<?php
// GENERATED CODE -- DO NOT EDIT!

namespace JournalClientInterface;

/**
 */
class JournalClient extends \Grpc\BaseStub {

    /**
     * @param string $hostname hostname
     * @param array $opts channel options
     * @param \Grpc\Channel $channel (optional) re-use channel object
     */
    public function __construct($hostname, $opts, $channel = null) {
        parent::__construct($hostname, $opts, $channel);
    }

    /**
     * @param \JournalClientInterface\JournalRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     */
    public function Publish(\JournalClientInterface\JournalRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/JournalClientInterface.Journal/Publish',
        $argument,
        ['\JournalClientInterface\JournalResponse', 'decode'],
        $metadata, $options);
    }

    /**
     * @param \JournalClientInterface\JournalRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     */
    public function UnPublish(\JournalClientInterface\JournalRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/JournalClientInterface.Journal/UnPublish',
        $argument,
        ['\JournalClientInterface\JournalResponse', 'decode'],
        $metadata, $options);
    }

    /**
     * @param \JournalClientInterface\JournalApproveRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     */
    public function Approve(\JournalClientInterface\JournalApproveRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/JournalClientInterface.Journal/Approve',
        $argument,
        ['\JournalClientInterface\JournalResponse', 'decode'],
        $metadata, $options);
    }

    /**
     * @param \JournalClientInterface\JournalRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     */
    public function Destroy(\JournalClientInterface\JournalRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/JournalClientInterface.Journal/Destroy',
        $argument,
        ['\JournalClientInterface\JournalResponse', 'decode'],
        $metadata, $options);
    }

    /**
     * @param \JournalClientInterface\JournalCreateRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     */
    public function Create(\JournalClientInterface\JournalCreateRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/JournalClientInterface.Journal/Create',
        $argument,
        ['\JournalClientInterface\JournalResponse', 'decode'],
        $metadata, $options);
    }

    /**
     * @param \JournalClientInterface\JournalEditRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     */
    public function Edit(\JournalClientInterface\JournalEditRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/JournalClientInterface.Journal/Edit',
        $argument,
        ['\JournalClientInterface\JournalResponse', 'decode'],
        $metadata, $options);
    }

}
