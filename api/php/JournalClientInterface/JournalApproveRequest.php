<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: journal.proto

namespace JournalClientInterface;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>JournalClientInterface.JournalApproveRequest</code>
 */
class JournalApproveRequest extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>int64 id = 1;</code>
     */
    private $id = 0;
    /**
     * Generated from protobuf field <code>int64 approved_by = 2;</code>
     */
    private $approved_by = 0;
    /**
     * Generated from protobuf field <code>int32 status = 3;</code>
     */
    private $status = 0;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type int|string $id
     *     @type int|string $approved_by
     *     @type int $status
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Journal::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>int64 id = 1;</code>
     * @return int|string
     */
    public function getId()
    {
        return $this->id;
    }

    /**
     * Generated from protobuf field <code>int64 id = 1;</code>
     * @param int|string $var
     * @return $this
     */
    public function setId($var)
    {
        GPBUtil::checkInt64($var);
        $this->id = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>int64 approved_by = 2;</code>
     * @return int|string
     */
    public function getApprovedBy()
    {
        return $this->approved_by;
    }

    /**
     * Generated from protobuf field <code>int64 approved_by = 2;</code>
     * @param int|string $var
     * @return $this
     */
    public function setApprovedBy($var)
    {
        GPBUtil::checkInt64($var);
        $this->approved_by = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>int32 status = 3;</code>
     * @return int
     */
    public function getStatus()
    {
        return $this->status;
    }

    /**
     * Generated from protobuf field <code>int32 status = 3;</code>
     * @param int $var
     * @return $this
     */
    public function setStatus($var)
    {
        GPBUtil::checkInt32($var);
        $this->status = $var;

        return $this;
    }

}

