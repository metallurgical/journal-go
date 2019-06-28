<?php
// GENERATED CODE -- DO NOT EDIT!

namespace RevisionClientInterface;

/**
 */
class RevisionClient extends \Grpc\BaseStub {

    /**
     * @param string $hostname hostname
     * @param array $opts channel options
     * @param \Grpc\Channel $channel (optional) re-use channel object
     */
    public function __construct($hostname, $opts, $channel = null) {
        parent::__construct($hostname, $opts, $channel);
    }

    /**
     * @param \RevisionClientInterface\RevisionRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     */
    public function Create(\RevisionClientInterface\RevisionRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/RevisionClientInterface.Revision/Create',
        $argument,
        ['\RevisionClientInterface\RevisionResponse', 'decode'],
        $metadata, $options);
    }

}
