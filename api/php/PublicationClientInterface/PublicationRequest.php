<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: publication.proto

namespace PublicationClientInterface;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>PublicationClientInterface.PublicationRequest</code>
 */
class PublicationRequest extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>string volume = 1;</code>
     */
    private $volume = '';
    /**
     * Generated from protobuf field <code>int64 no = 2;</code>
     */
    private $no = 0;
    /**
     * Generated from protobuf field <code>int64 created_by = 3;</code>
     */
    private $created_by = 0;
    /**
     * Generated from protobuf field <code>string published_at = 4;</code>
     */
    private $published_at = '';
    /**
     * Generated from protobuf field <code>int32 status = 5;</code>
     */
    private $status = 0;
    /**
     * Generated from protobuf field <code>int64 id = 6;</code>
     */
    private $id = 0;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $volume
     *     @type int|string $no
     *     @type int|string $created_by
     *     @type string $published_at
     *     @type int $status
     *     @type int|string $id
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Publication::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>string volume = 1;</code>
     * @return string
     */
    public function getVolume()
    {
        return $this->volume;
    }

    /**
     * Generated from protobuf field <code>string volume = 1;</code>
     * @param string $var
     * @return $this
     */
    public function setVolume($var)
    {
        GPBUtil::checkString($var, True);
        $this->volume = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>int64 no = 2;</code>
     * @return int|string
     */
    public function getNo()
    {
        return $this->no;
    }

    /**
     * Generated from protobuf field <code>int64 no = 2;</code>
     * @param int|string $var
     * @return $this
     */
    public function setNo($var)
    {
        GPBUtil::checkInt64($var);
        $this->no = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>int64 created_by = 3;</code>
     * @return int|string
     */
    public function getCreatedBy()
    {
        return $this->created_by;
    }

    /**
     * Generated from protobuf field <code>int64 created_by = 3;</code>
     * @param int|string $var
     * @return $this
     */
    public function setCreatedBy($var)
    {
        GPBUtil::checkInt64($var);
        $this->created_by = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string published_at = 4;</code>
     * @return string
     */
    public function getPublishedAt()
    {
        return $this->published_at;
    }

    /**
     * Generated from protobuf field <code>string published_at = 4;</code>
     * @param string $var
     * @return $this
     */
    public function setPublishedAt($var)
    {
        GPBUtil::checkString($var, True);
        $this->published_at = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>int32 status = 5;</code>
     * @return int
     */
    public function getStatus()
    {
        return $this->status;
    }

    /**
     * Generated from protobuf field <code>int32 status = 5;</code>
     * @param int $var
     * @return $this
     */
    public function setStatus($var)
    {
        GPBUtil::checkInt32($var);
        $this->status = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>int64 id = 6;</code>
     * @return int|string
     */
    public function getId()
    {
        return $this->id;
    }

    /**
     * Generated from protobuf field <code>int64 id = 6;</code>
     * @param int|string $var
     * @return $this
     */
    public function setId($var)
    {
        GPBUtil::checkInt64($var);
        $this->id = $var;

        return $this;
    }

}
