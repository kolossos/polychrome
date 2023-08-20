#pragma once
#include <juce_audio_utils/juce_audio_utils.h>

#include <memory>

#include "error.h"
#include "filter.h"
#include "processor.h"

namespace beak
{
class Engine
{
 public:
  struct Config
  {
    explicit Config() :
      m_deviceName(defaultDevice),
      m_inputs(defaultInputs),
      m_outputs(defaultOutputs),
      m_sampleRate(defaultSampleRate)
    {
    }

    Config WithDeviceName(juce::String const &deviceName)
    {
      auto retval = *this;
      retval.m_deviceName = deviceName.isEmpty() ? defaultDevice : deviceName;
      return retval;
    }
    Config WithInputs(int inputs)
    {
      auto retval = *this;
      retval.m_inputs = inputs;
      return retval;
    }
    Config WithOutputs(int outputs)
    {
      auto retval = *this;
      retval.m_outputs = outputs;
      return retval;
    }
    Config WithSampleRate(int sampleRate)
    {
      auto retval = *this;
      retval.m_sampleRate = sampleRate == 0 ? defaultSampleRate : sampleRate;
      return retval;
    }
    juce::String deviceName() const { return m_deviceName; }
    int inputs() const { return m_inputs; }
    int outputs() const { return m_outputs; }
    int sampleRate() const { return m_sampleRate; }

   private:
    juce::String m_deviceName;
    int m_inputs;
    int m_outputs;
    int m_sampleRate;

   public:
    static constexpr const char *defaultDevice = "MacBook Pro Speakers";
    static constexpr uint32_t defaultInputs = 2;
    static constexpr uint32_t defaultOutputs = 2;
    static constexpr int defaultSampleRate = 44100;
  };

 public:
  Engine();
  virtual ~Engine();
  Engine(Engine &&) = delete;
  Engine &operator=(Engine &&) = delete;

 public:
  [[nodiscard]] Error configure(Config const &config);
  [[nodiscard]] virtual Error playSound(const juce::File &file, int channel);
  [[nodiscard]] virtual Error stopPlayback(int channel);
  [[nodiscard]] virtual Error playSynth(const juce::MidiMessage &msg, int maxDurationMs = 1000);
  [[nodiscard]] virtual Error configureSynth(int channel, synth::Oscillator::Parameters &osc,
                                             const juce::ADSR::Parameters &adsr,
                                             const synth::Filter::Parameters &filter,
                                             const juce::ADSR::Parameters &filterAdsr);

 private:
  [[nodiscard]] virtual Error configureDeviceManager(Config const &config);
  [[nodiscard]] virtual Error configureGraph(Config const &config);

 protected:
  juce::AudioDeviceManager m_deviceManager;
  std::unique_ptr<juce::AudioProcessorGraph> m_mainProcessor;
  std::unique_ptr<juce::AudioProcessorPlayer> m_player;
  juce::AudioProcessorGraph::Node::Ptr m_audioOutputNode;
  std::vector<juce::AudioProcessorGraph::Node::Ptr> m_playerNodes;
  std::vector<juce::AudioProcessorGraph::Node::Ptr> m_synthNodes;

 private:
  JUCE_DECLARE_NON_COPYABLE_WITH_LEAK_DETECTOR(Engine)
};
}  // namespace beak
