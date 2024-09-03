// @generated by protoc-gen-es v2.0.0 with parameter "target=ts,import_extension=.ts"
// @generated from file protobuf/model/message/v1/message.proto (package protobuf.model.message.v1, syntax proto3)
/* eslint-disable */

import type { GenFile, GenMessage } from "@bufbuild/protobuf/codegenv1";
import { fileDesc, messageDesc } from "@bufbuild/protobuf/codegenv1";
import type { Message as Message$1 } from "@bufbuild/protobuf";

/**
 * Describes the file protobuf/model/message/v1/message.proto.
 */
export const file_protobuf_model_message_v1_message: GenFile = /*@__PURE__*/
  fileDesc("Cidwcm90b2J1Zi9tb2RlbC9tZXNzYWdlL3YxL21lc3NhZ2UucHJvdG8SGXByb3RvYnVmLm1vZGVsLm1lc3NhZ2UudjEiUwoNTWVzc2FnZVNlbmRlchIPCgd1c2VyX2lkGAEgASgJEgwKBG5hbWUYAiABKAkSEQoJbGFzdF9uYW1lGAMgASgJEhAKCHVzZXJuYW1lGAQgASgJIqQCCgdNZXNzYWdlEhIKCm1lc3NhZ2VfaWQYASABKAkSOAoGc2VuZGVyGAIgASgLMigucHJvdG9idWYubW9kZWwubWVzc2FnZS52MS5NZXNzYWdlU2VuZGVyEhIKCmNyZWF0ZWRfYXQYAyABKAMSDgoGZWRpdGVkGAQgASgIEgwKBHNlZW4YBSABKAgSDAoEdHlwZRgGIAEoCRI+Cgx0ZXh0X21lc3NhZ2UYByABKAsyJi5wcm90b2J1Zi5tb2RlbC5tZXNzYWdlLnYxLlRleHRNZXNzYWdlSAASQAoNbGFiZWxfbWVzc2FnZRgIIAEoCzInLnByb3RvYnVmLm1vZGVsLm1lc3NhZ2UudjEuTGFiZWxNZXNzYWdlSABCCQoHcGF5bG9hZCIbCgtUZXh0TWVzc2FnZRIMCgR0ZXh0GAEgASgJIhwKDExhYmVsTWVzc2FnZRIMCgR0ZXh0GAEgASgJQocCCh1jb20ucHJvdG9idWYubW9kZWwubWVzc2FnZS52MUIMTWVzc2FnZVByb3RvUAFaUWdpdGh1Yi5jb20va2F2a2Fjby9LYXZrYS1Db3JlL3Byb3RvYnVmL2dlbi9nby9wcm90b2J1Zi9tb2RlbC9tZXNzYWdlL3YxO21lc3NhZ2V2MaICA1BNTaoCGVByb3RvYnVmLk1vZGVsLk1lc3NhZ2UuVjHKAhlQcm90b2J1ZlxNb2RlbFxNZXNzYWdlXFYx4gIlUHJvdG9idWZcTW9kZWxcTWVzc2FnZVxWMVxHUEJNZXRhZGF0YeoCHFByb3RvYnVmOjpNb2RlbDo6TWVzc2FnZTo6VjFiBnByb3RvMw");

/**
 * @generated from message protobuf.model.message.v1.MessageSender
 */
export type MessageSender = Message$1<"protobuf.model.message.v1.MessageSender"> & {
  /**
   * @generated from field: string user_id = 1;
   */
  userId: string;

  /**
   * @generated from field: string name = 2;
   */
  name: string;

  /**
   * @generated from field: string last_name = 3;
   */
  lastName: string;

  /**
   * @generated from field: string username = 4;
   */
  username: string;
};

/**
 * Describes the message protobuf.model.message.v1.MessageSender.
 * Use `create(MessageSenderSchema)` to create a new message.
 */
export const MessageSenderSchema: GenMessage<MessageSender> = /*@__PURE__*/
  messageDesc(file_protobuf_model_message_v1_message, 0);

/**
 * @generated from message protobuf.model.message.v1.Message
 */
export type Message = Message$1<"protobuf.model.message.v1.Message"> & {
  /**
   * @generated from field: string message_id = 1;
   */
  messageId: string;

  /**
   * @generated from field: protobuf.model.message.v1.MessageSender sender = 2;
   */
  sender?: MessageSender;

  /**
   * @generated from field: int64 created_at = 3;
   */
  createdAt: bigint;

  /**
   * @generated from field: bool edited = 4;
   */
  edited: boolean;

  /**
   * @generated from field: bool seen = 5;
   */
  seen: boolean;

  /**
   * @generated from field: string type = 6;
   */
  type: string;

  /**
   * @generated from oneof protobuf.model.message.v1.Message.payload
   */
  payload: {
    /**
     * @generated from field: protobuf.model.message.v1.TextMessage text_message = 7;
     */
    value: TextMessage;
    case: "textMessage";
  } | {
    /**
     * @generated from field: protobuf.model.message.v1.LabelMessage label_message = 8;
     */
    value: LabelMessage;
    case: "labelMessage";
  } | { case: undefined; value?: undefined };
};

/**
 * Describes the message protobuf.model.message.v1.Message.
 * Use `create(MessageSchema)` to create a new message.
 */
export const MessageSchema: GenMessage<Message> = /*@__PURE__*/
  messageDesc(file_protobuf_model_message_v1_message, 1);

/**
 * @generated from message protobuf.model.message.v1.TextMessage
 */
export type TextMessage = Message$1<"protobuf.model.message.v1.TextMessage"> & {
  /**
   * @generated from field: string text = 1;
   */
  text: string;
};

/**
 * Describes the message protobuf.model.message.v1.TextMessage.
 * Use `create(TextMessageSchema)` to create a new message.
 */
export const TextMessageSchema: GenMessage<TextMessage> = /*@__PURE__*/
  messageDesc(file_protobuf_model_message_v1_message, 2);

/**
 * @generated from message protobuf.model.message.v1.LabelMessage
 */
export type LabelMessage = Message$1<"protobuf.model.message.v1.LabelMessage"> & {
  /**
   * @generated from field: string text = 1;
   */
  text: string;
};

/**
 * Describes the message protobuf.model.message.v1.LabelMessage.
 * Use `create(LabelMessageSchema)` to create a new message.
 */
export const LabelMessageSchema: GenMessage<LabelMessage> = /*@__PURE__*/
  messageDesc(file_protobuf_model_message_v1_message, 3);

