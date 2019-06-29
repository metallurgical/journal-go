<?php
// GENERATED CODE -- DO NOT EDIT!

namespace PublicationClientInterface;

/**
 */
class PublicationClient extends \Grpc\BaseStub {

    /**
     * @param string $hostname hostname
     * @param array $opts channel options
     * @param \Grpc\Channel $channel (optional) re-use channel object
     */
    public function __construct($hostname, $opts, $channel = null) {
        parent::__construct($hostname, $opts, $channel);
    }

    /**
     * @param \PublicationClientInterface\PublicationRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     */
    public function Create(\PublicationClientInterface\PublicationRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/PublicationClientInterface.Publication/Create',
        $argument,
        ['\PublicationClientInterface\PublicationResponse', 'decode'],
        $metadata, $options);
    }

    /**
     * @param \PublicationClientInterface\PublicationRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     */
    public function Update(\PublicationClientInterface\PublicationRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/PublicationClientInterface.Publication/Update',
        $argument,
        ['\PublicationClientInterface\PublicationResponse', 'decode'],
        $metadata, $options);
    }

    /**
     * @param \PublicationClientInterface\PublicationRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     */
    public function Delete(\PublicationClientInterface\PublicationRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/PublicationClientInterface.Publication/Delete',
        $argument,
        ['\PublicationClientInterface\PublicationResponse', 'decode'],
        $metadata, $options);
    }

    /**
     * @param \PublicationClientInterface\PublicationRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     */
    public function SetStatus(\PublicationClientInterface\PublicationRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/PublicationClientInterface.Publication/SetStatus',
        $argument,
        ['\PublicationClientInterface\PublicationResponse', 'decode'],
        $metadata, $options);
    }

}
