// @generated by protoc-gen-es v2.0.0 with parameter "target=ts,import_extension=.ts"
// @generated from file protobuf/search/v1/search.proto (package protobuf.search.v1, syntax proto3)
/* eslint-disable */

import type { GenFile, GenMessage, GenService } from "@bufbuild/protobuf/codegenv1";
import { fileDesc, messageDesc, serviceDesc } from "@bufbuild/protobuf/codegenv1";
import type { Chat } from "../../model/chat/v1/chat_pb.ts";
import { file_protobuf_model_chat_v1_chat } from "../../model/chat/v1/chat_pb.ts";
import type { User } from "../../model/user/v1/user_pb.ts";
import { file_protobuf_model_user_v1_user } from "../../model/user/v1/user_pb.ts";
import type { Message } from "@bufbuild/protobuf";

/**
 * Describes the file protobuf/search/v1/search.proto.
 */
export const file_protobuf_search_v1_search: GenFile = /*@__PURE__*/
  fileDesc("Ch9wcm90b2J1Zi9zZWFyY2gvdjEvc2VhcmNoLnByb3RvEhJwcm90b2J1Zi5zZWFyY2gudjEiHgoNU2VhcmNoUmVxdWVzdBINCgVpbnB1dBgBIAEoCSK7AQoOU2VhcmNoUmVzcG9uc2USPwoGcmVzdWx0GAEgASgLMi8ucHJvdG9idWYuc2VhcmNoLnYxLlNlYXJjaFJlc3BvbnNlLlNlYXJjaFJlc3VsdBpoCgxTZWFyY2hSZXN1bHQSKwoFY2hhdHMYASADKAsyHC5wcm90b2J1Zi5tb2RlbC5jaGF0LnYxLkNoYXQSKwoFdXNlcnMYAiADKAsyHC5wcm90b2J1Zi5tb2RlbC51c2VyLnYxLlVzZXIyYgoNU2VhcmNoU2VydmljZRJRCgZTZWFyY2gSIS5wcm90b2J1Zi5zZWFyY2gudjEuU2VhcmNoUmVxdWVzdBoiLnByb3RvYnVmLnNlYXJjaC52MS5TZWFyY2hSZXNwb25zZSIAQtoBChZjb20ucHJvdG9idWYuc2VhcmNoLnYxQgtTZWFyY2hQcm90b1ABWklnaXRodWIuY29tL2thdmthY28vS2F2a2EtQ29yZS9wcm90b2J1Zi9nZW4vZ28vcHJvdG9idWYvc2VhcmNoL3YxO3NlYXJjaHYxogIDUFNYqgISUHJvdG9idWYuU2VhcmNoLlYxygISUHJvdG9idWZcU2VhcmNoXFYx4gIeUHJvdG9idWZcU2VhcmNoXFYxXEdQQk1ldGFkYXRh6gIUUHJvdG9idWY6OlNlYXJjaDo6VjFiBnByb3RvMw", [file_protobuf_model_chat_v1_chat, file_protobuf_model_user_v1_user]);

/**
 * @generated from message protobuf.search.v1.SearchRequest
 */
export type SearchRequest = Message<"protobuf.search.v1.SearchRequest"> & {
  /**
   * @generated from field: string input = 1;
   */
  input: string;
};

/**
 * Describes the message protobuf.search.v1.SearchRequest.
 * Use `create(SearchRequestSchema)` to create a new message.
 */
export const SearchRequestSchema: GenMessage<SearchRequest> = /*@__PURE__*/
  messageDesc(file_protobuf_search_v1_search, 0);

/**
 * @generated from message protobuf.search.v1.SearchResponse
 */
export type SearchResponse = Message<"protobuf.search.v1.SearchResponse"> & {
  /**
   * @generated from field: protobuf.search.v1.SearchResponse.SearchResult result = 1;
   */
  result?: SearchResponse_SearchResult;
};

/**
 * Describes the message protobuf.search.v1.SearchResponse.
 * Use `create(SearchResponseSchema)` to create a new message.
 */
export const SearchResponseSchema: GenMessage<SearchResponse> = /*@__PURE__*/
  messageDesc(file_protobuf_search_v1_search, 1);

/**
 * @generated from message protobuf.search.v1.SearchResponse.SearchResult
 */
export type SearchResponse_SearchResult = Message<"protobuf.search.v1.SearchResponse.SearchResult"> & {
  /**
   * @generated from field: repeated protobuf.model.chat.v1.Chat chats = 1;
   */
  chats: Chat[];

  /**
   * @generated from field: repeated protobuf.model.user.v1.User users = 2;
   */
  users: User[];
};

/**
 * Describes the message protobuf.search.v1.SearchResponse.SearchResult.
 * Use `create(SearchResponse_SearchResultSchema)` to create a new message.
 */
export const SearchResponse_SearchResultSchema: GenMessage<SearchResponse_SearchResult> = /*@__PURE__*/
  messageDesc(file_protobuf_search_v1_search, 1, 0);

/**
 * @generated from service protobuf.search.v1.SearchService
 */
export const SearchService: GenService<{
  /**
   * @generated from rpc protobuf.search.v1.SearchService.Search
   */
  search: {
    methodKind: "unary";
    input: typeof SearchRequestSchema;
    output: typeof SearchResponseSchema;
  },
}> = /*@__PURE__*/
  serviceDesc(file_protobuf_search_v1_search, 0);

