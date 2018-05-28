// +build darwin,mrb linux,mrb

package rpc_cached

import (
	"log"
	"strings"

	"github.com/anycable/anycable-go/mrb"
	"github.com/anycable/anycable-go/node"
	mruby "github.com/mitchellh/go-mruby"
)

// MCallResult is an interface to decode mruby call
type MCallResult struct {
	Streams        []string `mruby:"streams"`
	StopAllStreams bool     `mruby:"stop_all_streams"`
	Transmissions  []string `mruby:"transmissions"`
}

// MCache is a cache of mruby compiled channels methods
type MCache struct {
	engine *mrb.Engine
	store  map[string]map[string]*MAction
}

var (
	// MrbCache is a global channels methods cache
	MrbCache *MCache
)

// InitCache creates global cache
func InitCache() {
	if MrbCache != nil {
		return
	}

	var err error
	MrbCache, err = NewMCache(mrb.DefaultEngine())

	if err != nil {
		log.Fatalf("Failed to initialize MrbCache: %s", err)
	}
}

// NewMCache builds a new cache struct for mruby engine
func NewMCache(mengine *mrb.Engine) (*MCache, error) {
	// Build base channel class
	err := mengine.LoadString(
		`
	module AnyCable
		class Channel
			class Result
				attr_reader :streams, :transmissions
				attr_accessor :stop_all_streams

				def initialize
					@streams = []
					@transmissions = []
					@stop_all_streams = false
				end

				def to_gomruby
					{
						streams: streams,
						transmissions: transmissions,
						stop_all_streams: stop_all_streams
					}
				end
			end

			class << self
				attr_reader :identifier

				def identify(channel)
					@identifier = { channel: channel }.to_json
				end

				def perform(json_data)
					data = JSON.parse(json_data)
					new.tap do |channel|
						channel.send(data.fetch('action'), data)
					end.result
				end
			end

			attr_reader :result

			def initialize
				@result = Result.new
			end

			def transmit(data)
				result.transmissions << {
					identifier: self.class.identifier,
					message: data
				}.to_json
			end
		end
	end
	`,
	)

	if err != nil {
		return nil, err
	}

	return &MCache{
		engine: mengine,
		store:  make(map[string]map[string]*MAction),
	}, nil
}

// MAction is a signle cached channel method
type MAction struct {
	compiled *mruby.MrbValue
}

// NewMAction compiles a channel method within mruby VM
func NewMAction(engine *mrb.Engine, channel string, source string) (*MAction, error) {
	var buf strings.Builder

	channelClass := "CachedChannel_" + channel

	buf.WriteString(
		"class " + channelClass + " < AnyCable::Channel\n",
	)
	buf.WriteString("identify \"" + channel + "\"\n")
	buf.WriteString(source + "\n")
	buf.WriteString("end\n")

	err := engine.LoadString(buf.String())

	if err != nil {
		return nil, err
	}

	mchannel := engine.VM.Class(channelClass, nil)

	mchannelValue := mchannel.MrbValue(engine.VM)

	return &MAction{compiled: mchannelValue}, nil
}

// Perform executes action within mruby
func (m *MAction) Perform(data string) (*node.CommandResult, error) {
	result, err := m.compiled.Call("perform", mruby.String(data))

	if err != nil {
		return nil, err
	}

	decoded := MCallResult{}

	err = mruby.Decode(&decoded, result)

	if err != nil {
		return nil, err
	}

	res := &node.CommandResult{
		Transmissions:  decoded.Transmissions,
		StopAllStreams: decoded.StopAllStreams,
		Streams:        decoded.Streams,
	}

	return res, nil
}
