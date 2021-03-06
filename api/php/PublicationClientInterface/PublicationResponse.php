<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: publication.proto

namespace PublicationClientInterface;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>PublicationClientInterface.PublicationResponse</code>
 */
class PublicationResponse extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>string flag = 1;</code>
     */
    private $flag = '';
    /**
     * Generated from protobuf field <code>string message = 2;</code>
     */
    private $message = '';
    /**
     * Generated from protobuf field <code>string Publication = 3;</code>
     */
    private $Publication = '';

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $flag
     *     @type string $message
     *     @type string $Publication
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Publication::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>string flag = 1;</code>
     * @return string
     */
    public function getFlag()
    {
        return $this->flag;
    }

    /**
     * Generated from protobuf field <code>string flag = 1;</code>
     * @param string $var
     * @return $this
     */
    public function setFlag($var)
    {
        GPBUtil::checkString($var, True);
        $this->flag = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string message = 2;</code>
     * @return string
     */
    public function getMessage()
    {
        return $this->message;
    }

    /**
     * Generated from protobuf field <code>string message = 2;</code>
     * @param string $var
     * @return $this
     */
    public function setMessage($var)
    {
        GPBUtil::checkString($var, True);
        $this->message = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string Publication = 3;</code>
     * @return string
     */
    public function getPublication()
    {
        return $this->Publication;
    }

    /**
     * Generated from protobuf field <code>string Publication = 3;</code>
     * @param string $var
     * @return $this
     */
    public function setPublication($var)
    {
        GPBUtil::checkString($var, True);
        $this->Publication = $var;

        return $this;
    }

}

