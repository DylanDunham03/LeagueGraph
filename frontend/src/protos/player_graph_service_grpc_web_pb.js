/**
 * @fileoverview gRPC-Web generated client stub for playergraph
 * @enhanceable
 * @public
 */

// Code generated by protoc-gen-grpc-web. DO NOT EDIT.
// versions:
// 	protoc-gen-grpc-web v1.5.0
// 	protoc              v5.27.0
// source: player_graph_service.proto


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');


var google_protobuf_empty_pb = require('google-protobuf/google/protobuf/empty_pb.js')

var google_api_annotations_pb = require('./google/api/annotations_pb.js')
const proto = {};
proto.playergraph = require('./player_graph_service_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.playergraph.PlayerGraphServiceClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname.replace(/\/+$/, '');

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.playergraph.PlayerGraphServicePromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname.replace(/\/+$/, '');

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.google.protobuf.Empty,
 *   !proto.playergraph.PlayerResponse>}
 */
const methodDescriptor_PlayerGraphService_GetPlayerData = new grpc.web.MethodDescriptor(
  '/playergraph.PlayerGraphService/GetPlayerData',
  grpc.web.MethodType.UNARY,
  google_protobuf_empty_pb.Empty,
  proto.playergraph.PlayerResponse,
  /**
   * @param {!proto.google.protobuf.Empty} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.playergraph.PlayerResponse.deserializeBinary
);


/**
 * @param {!proto.google.protobuf.Empty} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.playergraph.PlayerResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.playergraph.PlayerResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.playergraph.PlayerGraphServiceClient.prototype.getPlayerData =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/playergraph.PlayerGraphService/GetPlayerData',
      request,
      metadata || {},
      methodDescriptor_PlayerGraphService_GetPlayerData,
      callback);
};


/**
 * @param {!proto.google.protobuf.Empty} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.playergraph.PlayerResponse>}
 *     Promise that resolves to the response
 */
proto.playergraph.PlayerGraphServicePromiseClient.prototype.getPlayerData =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/playergraph.PlayerGraphService/GetPlayerData',
      request,
      metadata || {},
      methodDescriptor_PlayerGraphService_GetPlayerData);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.playergraph.GraphRequest,
 *   !proto.playergraph.GraphResponse>}
 */
const methodDescriptor_PlayerGraphService_GetPlayerGraph = new grpc.web.MethodDescriptor(
  '/playergraph.PlayerGraphService/GetPlayerGraph',
  grpc.web.MethodType.UNARY,
  proto.playergraph.GraphRequest,
  proto.playergraph.GraphResponse,
  /**
   * @param {!proto.playergraph.GraphRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.playergraph.GraphResponse.deserializeBinary
);


/**
 * @param {!proto.playergraph.GraphRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.playergraph.GraphResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.playergraph.GraphResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.playergraph.PlayerGraphServiceClient.prototype.getPlayerGraph =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/playergraph.PlayerGraphService/GetPlayerGraph',
      request,
      metadata || {},
      methodDescriptor_PlayerGraphService_GetPlayerGraph,
      callback);
};


/**
 * @param {!proto.playergraph.GraphRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.playergraph.GraphResponse>}
 *     Promise that resolves to the response
 */
proto.playergraph.PlayerGraphServicePromiseClient.prototype.getPlayerGraph =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/playergraph.PlayerGraphService/GetPlayerGraph',
      request,
      metadata || {},
      methodDescriptor_PlayerGraphService_GetPlayerGraph);
};


module.exports = proto.playergraph;

