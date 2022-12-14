// source: v1/client.proto
/**
 * @fileoverview
 * @enhanceable
 * @suppress {missingRequire} reports error on implicit type usages.
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!
/* eslint-disable */
// @ts-nocheck

var jspb = require('google-protobuf');
var goog = jspb;
var global = Function('return this')();

var google_protobuf_timestamp_pb = require('google-protobuf/google/protobuf/timestamp_pb.js');
goog.object.extend(proto, google_protobuf_timestamp_pb);
var google_protobuf_any_pb = require('google-protobuf/google/protobuf/any_pb.js');
goog.object.extend(proto, google_protobuf_any_pb);
var v1_application_type_pb = require('../v1/application_type_pb.js');
goog.object.extend(proto, v1_application_type_pb);
var v1_client_approval_pb = require('../v1/client_approval_pb.js');
goog.object.extend(proto, v1_client_approval_pb);
var v1_grant_type_pb = require('../v1/grant_type_pb.js');
goog.object.extend(proto, v1_grant_type_pb);
var v1_public_keys_configuration_pb = require('../v1/public_keys_configuration_pb.js');
goog.object.extend(proto, v1_public_keys_configuration_pb);
var v1_refresh_token_rotation_type_pb = require('../v1/refresh_token_rotation_type_pb.js');
goog.object.extend(proto, v1_refresh_token_rotation_type_pb);
var v1_response_type_pb = require('../v1/response_type_pb.js');
goog.object.extend(proto, v1_response_type_pb);
var v1_token_endpoint_auth_method_pb = require('../v1/token_endpoint_auth_method_pb.js');
goog.object.extend(proto, v1_token_endpoint_auth_method_pb);
goog.exportSymbol('proto.v1.Client', null, global);
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.v1.Client = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.v1.Client.repeatedFields_, null);
};
goog.inherits(proto.v1.Client, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.v1.Client.displayName = 'proto.v1.Client';
}

/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.v1.Client.repeatedFields_ = [3,4,6,8,14,15,16,18];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.v1.Client.prototype.toObject = function(opt_includeInstance) {
  return proto.v1.Client.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.v1.Client} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.v1.Client.toObject = function(includeInstance, msg) {
  var f, obj = {
    id: jspb.Message.getFieldWithDefault(msg, 1, ""),
    userId: jspb.Message.getFieldWithDefault(msg, 2, ""),
    contactsList: (f = jspb.Message.getRepeatedField(msg, 3)) == null ? undefined : f,
    clientApprovalsList: jspb.Message.toObjectList(msg.getClientApprovalsList(),
    v1_client_approval_pb.ClientApproval.toObject, includeInstance),
    description: jspb.Message.getFieldWithDefault(msg, 5, ""),
    grantTypesList: (f = jspb.Message.getRepeatedField(msg, 6)) == null ? undefined : f,
    isFirstParty: jspb.Message.getBooleanFieldWithDefault(msg, 7, false),
    jwksList: jspb.Message.toObjectList(msg.getJwksList(),
    google_protobuf_any_pb.Any.toObject, includeInstance),
    jwksUri: jspb.Message.getFieldWithDefault(msg, 9, ""),
    logoUri: jspb.Message.getFieldWithDefault(msg, 10, ""),
    name: jspb.Message.getFieldWithDefault(msg, 11, ""),
    policyUri: jspb.Message.getFieldWithDefault(msg, 12, ""),
    publicKeysConfiguration: jspb.Message.getFieldWithDefault(msg, 13, 0),
    redirectUrisList: (f = jspb.Message.getRepeatedField(msg, 14)) == null ? undefined : f,
    refreshTokenRotationTypeList: (f = jspb.Message.getRepeatedField(msg, 15)) == null ? undefined : f,
    responseTypesList: (f = jspb.Message.getRepeatedField(msg, 16)) == null ? undefined : f,
    secret: jspb.Message.getFieldWithDefault(msg, 17, ""),
    scopeList: (f = jspb.Message.getRepeatedField(msg, 18)) == null ? undefined : f,
    softwareId: jspb.Message.getFieldWithDefault(msg, 19, ""),
    softwareVersion: jspb.Message.getFieldWithDefault(msg, 20, ""),
    type: jspb.Message.getFieldWithDefault(msg, 21, 0),
    tosUri: jspb.Message.getFieldWithDefault(msg, 22, ""),
    tokenEndpointAuthMethod: jspb.Message.getFieldWithDefault(msg, 23, 0),
    uri: jspb.Message.getFieldWithDefault(msg, 24, ""),
    createdAt: (f = msg.getCreatedAt()) && google_protobuf_timestamp_pb.Timestamp.toObject(includeInstance, f),
    updatedAt: (f = msg.getUpdatedAt()) && google_protobuf_timestamp_pb.Timestamp.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.v1.Client}
 */
proto.v1.Client.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.v1.Client;
  return proto.v1.Client.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.v1.Client} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.v1.Client}
 */
proto.v1.Client.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setId(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setUserId(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.addContacts(value);
      break;
    case 4:
      var value = new v1_client_approval_pb.ClientApproval;
      reader.readMessage(value,v1_client_approval_pb.ClientApproval.deserializeBinaryFromReader);
      msg.addClientApprovals(value);
      break;
    case 5:
      var value = /** @type {string} */ (reader.readString());
      msg.setDescription(value);
      break;
    case 6:
      var values = /** @type {!Array<!proto.v1.GrantType>} */ (reader.isDelimited() ? reader.readPackedEnum() : [reader.readEnum()]);
      for (var i = 0; i < values.length; i++) {
        msg.addGrantTypes(values[i]);
      }
      break;
    case 7:
      var value = /** @type {boolean} */ (reader.readBool());
      msg.setIsFirstParty(value);
      break;
    case 8:
      var value = new google_protobuf_any_pb.Any;
      reader.readMessage(value,google_protobuf_any_pb.Any.deserializeBinaryFromReader);
      msg.addJwks(value);
      break;
    case 9:
      var value = /** @type {string} */ (reader.readString());
      msg.setJwksUri(value);
      break;
    case 10:
      var value = /** @type {string} */ (reader.readString());
      msg.setLogoUri(value);
      break;
    case 11:
      var value = /** @type {string} */ (reader.readString());
      msg.setName(value);
      break;
    case 12:
      var value = /** @type {string} */ (reader.readString());
      msg.setPolicyUri(value);
      break;
    case 13:
      var value = /** @type {!proto.v1.PublicKeysConfiguration} */ (reader.readEnum());
      msg.setPublicKeysConfiguration(value);
      break;
    case 14:
      var value = /** @type {string} */ (reader.readString());
      msg.addRedirectUris(value);
      break;
    case 15:
      var values = /** @type {!Array<!proto.v1.RefreshTokenRotationType>} */ (reader.isDelimited() ? reader.readPackedEnum() : [reader.readEnum()]);
      for (var i = 0; i < values.length; i++) {
        msg.addRefreshTokenRotationType(values[i]);
      }
      break;
    case 16:
      var values = /** @type {!Array<!proto.v1.ResponseType>} */ (reader.isDelimited() ? reader.readPackedEnum() : [reader.readEnum()]);
      for (var i = 0; i < values.length; i++) {
        msg.addResponseTypes(values[i]);
      }
      break;
    case 17:
      var value = /** @type {string} */ (reader.readString());
      msg.setSecret(value);
      break;
    case 18:
      var value = /** @type {string} */ (reader.readString());
      msg.addScope(value);
      break;
    case 19:
      var value = /** @type {string} */ (reader.readString());
      msg.setSoftwareId(value);
      break;
    case 20:
      var value = /** @type {string} */ (reader.readString());
      msg.setSoftwareVersion(value);
      break;
    case 21:
      var value = /** @type {!proto.v1.ApplicationType} */ (reader.readEnum());
      msg.setType(value);
      break;
    case 22:
      var value = /** @type {string} */ (reader.readString());
      msg.setTosUri(value);
      break;
    case 23:
      var value = /** @type {!proto.v1.TokenEndpointAuthMethod} */ (reader.readEnum());
      msg.setTokenEndpointAuthMethod(value);
      break;
    case 24:
      var value = /** @type {string} */ (reader.readString());
      msg.setUri(value);
      break;
    case 25:
      var value = new google_protobuf_timestamp_pb.Timestamp;
      reader.readMessage(value,google_protobuf_timestamp_pb.Timestamp.deserializeBinaryFromReader);
      msg.setCreatedAt(value);
      break;
    case 26:
      var value = new google_protobuf_timestamp_pb.Timestamp;
      reader.readMessage(value,google_protobuf_timestamp_pb.Timestamp.deserializeBinaryFromReader);
      msg.setUpdatedAt(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.v1.Client.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.v1.Client.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.v1.Client} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.v1.Client.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getId();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getUserId();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getContactsList();
  if (f.length > 0) {
    writer.writeRepeatedString(
      3,
      f
    );
  }
  f = message.getClientApprovalsList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      4,
      f,
      v1_client_approval_pb.ClientApproval.serializeBinaryToWriter
    );
  }
  f = message.getDescription();
  if (f.length > 0) {
    writer.writeString(
      5,
      f
    );
  }
  f = message.getGrantTypesList();
  if (f.length > 0) {
    writer.writePackedEnum(
      6,
      f
    );
  }
  f = message.getIsFirstParty();
  if (f) {
    writer.writeBool(
      7,
      f
    );
  }
  f = message.getJwksList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      8,
      f,
      google_protobuf_any_pb.Any.serializeBinaryToWriter
    );
  }
  f = message.getJwksUri();
  if (f.length > 0) {
    writer.writeString(
      9,
      f
    );
  }
  f = message.getLogoUri();
  if (f.length > 0) {
    writer.writeString(
      10,
      f
    );
  }
  f = message.getName();
  if (f.length > 0) {
    writer.writeString(
      11,
      f
    );
  }
  f = message.getPolicyUri();
  if (f.length > 0) {
    writer.writeString(
      12,
      f
    );
  }
  f = message.getPublicKeysConfiguration();
  if (f !== 0.0) {
    writer.writeEnum(
      13,
      f
    );
  }
  f = message.getRedirectUrisList();
  if (f.length > 0) {
    writer.writeRepeatedString(
      14,
      f
    );
  }
  f = message.getRefreshTokenRotationTypeList();
  if (f.length > 0) {
    writer.writePackedEnum(
      15,
      f
    );
  }
  f = message.getResponseTypesList();
  if (f.length > 0) {
    writer.writePackedEnum(
      16,
      f
    );
  }
  f = message.getSecret();
  if (f.length > 0) {
    writer.writeString(
      17,
      f
    );
  }
  f = message.getScopeList();
  if (f.length > 0) {
    writer.writeRepeatedString(
      18,
      f
    );
  }
  f = message.getSoftwareId();
  if (f.length > 0) {
    writer.writeString(
      19,
      f
    );
  }
  f = message.getSoftwareVersion();
  if (f.length > 0) {
    writer.writeString(
      20,
      f
    );
  }
  f = message.getType();
  if (f !== 0.0) {
    writer.writeEnum(
      21,
      f
    );
  }
  f = message.getTosUri();
  if (f.length > 0) {
    writer.writeString(
      22,
      f
    );
  }
  f = message.getTokenEndpointAuthMethod();
  if (f !== 0.0) {
    writer.writeEnum(
      23,
      f
    );
  }
  f = message.getUri();
  if (f.length > 0) {
    writer.writeString(
      24,
      f
    );
  }
  f = message.getCreatedAt();
  if (f != null) {
    writer.writeMessage(
      25,
      f,
      google_protobuf_timestamp_pb.Timestamp.serializeBinaryToWriter
    );
  }
  f = message.getUpdatedAt();
  if (f != null) {
    writer.writeMessage(
      26,
      f,
      google_protobuf_timestamp_pb.Timestamp.serializeBinaryToWriter
    );
  }
};


/**
 * optional string id = 1;
 * @return {string}
 */
proto.v1.Client.prototype.getId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.v1.Client} returns this
 */
proto.v1.Client.prototype.setId = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string user_id = 2;
 * @return {string}
 */
proto.v1.Client.prototype.getUserId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.v1.Client} returns this
 */
proto.v1.Client.prototype.setUserId = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * repeated string contacts = 3;
 * @return {!Array<string>}
 */
proto.v1.Client.prototype.getContactsList = function() {
  return /** @type {!Array<string>} */ (jspb.Message.getRepeatedField(this, 3));
};


/**
 * @param {!Array<string>} value
 * @return {!proto.v1.Client} returns this
 */
proto.v1.Client.prototype.setContactsList = function(value) {
  return jspb.Message.setField(this, 3, value || []);
};


/**
 * @param {string} value
 * @param {number=} opt_index
 * @return {!proto.v1.Client} returns this
 */
proto.v1.Client.prototype.addContacts = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 3, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.v1.Client} returns this
 */
proto.v1.Client.prototype.clearContactsList = function() {
  return this.setContactsList([]);
};


/**
 * repeated ClientApproval client_approvals = 4;
 * @return {!Array<!proto.v1.ClientApproval>}
 */
proto.v1.Client.prototype.getClientApprovalsList = function() {
  return /** @type{!Array<!proto.v1.ClientApproval>} */ (
    jspb.Message.getRepeatedWrapperField(this, v1_client_approval_pb.ClientApproval, 4));
};


/**
 * @param {!Array<!proto.v1.ClientApproval>} value
 * @return {!proto.v1.Client} returns this
*/
proto.v1.Client.prototype.setClientApprovalsList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 4, value);
};


/**
 * @param {!proto.v1.ClientApproval=} opt_value
 * @param {number=} opt_index
 * @return {!proto.v1.ClientApproval}
 */
proto.v1.Client.prototype.addClientApprovals = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 4, opt_value, proto.v1.ClientApproval, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.v1.Client} returns this
 */
proto.v1.Client.prototype.clearClientApprovalsList = function() {
  return this.setClientApprovalsList([]);
};


/**
 * optional string description = 5;
 * @return {string}
 */
proto.v1.Client.prototype.getDescription = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 5, ""));
};


/**
 * @param {string} value
 * @return {!proto.v1.Client} returns this
 */
proto.v1.Client.prototype.setDescription = function(value) {
  return jspb.Message.setProto3StringField(this, 5, value);
};


/**
 * repeated GrantType grant_types = 6;
 * @return {!Array<!proto.v1.GrantType>}
 */
proto.v1.Client.prototype.getGrantTypesList = function() {
  return /** @type {!Array<!proto.v1.GrantType>} */ (jspb.Message.getRepeatedField(this, 6));
};


/**
 * @param {!Array<!proto.v1.GrantType>} value
 * @return {!proto.v1.Client} returns this
 */
proto.v1.Client.prototype.setGrantTypesList = function(value) {
  return jspb.Message.setField(this, 6, value || []);
};


/**
 * @param {!proto.v1.GrantType} value
 * @param {number=} opt_index
 * @return {!proto.v1.Client} returns this
 */
proto.v1.Client.prototype.addGrantTypes = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 6, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.v1.Client} returns this
 */
proto.v1.Client.prototype.clearGrantTypesList = function() {
  return this.setGrantTypesList([]);
};


/**
 * optional bool is_first_party = 7;
 * @return {boolean}
 */
proto.v1.Client.prototype.getIsFirstParty = function() {
  return /** @type {boolean} */ (jspb.Message.getBooleanFieldWithDefault(this, 7, false));
};


/**
 * @param {boolean} value
 * @return {!proto.v1.Client} returns this
 */
proto.v1.Client.prototype.setIsFirstParty = function(value) {
  return jspb.Message.setProto3BooleanField(this, 7, value);
};


/**
 * repeated google.protobuf.Any jwks = 8;
 * @return {!Array<!proto.google.protobuf.Any>}
 */
proto.v1.Client.prototype.getJwksList = function() {
  return /** @type{!Array<!proto.google.protobuf.Any>} */ (
    jspb.Message.getRepeatedWrapperField(this, google_protobuf_any_pb.Any, 8));
};


/**
 * @param {!Array<!proto.google.protobuf.Any>} value
 * @return {!proto.v1.Client} returns this
*/
proto.v1.Client.prototype.setJwksList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 8, value);
};


/**
 * @param {!proto.google.protobuf.Any=} opt_value
 * @param {number=} opt_index
 * @return {!proto.google.protobuf.Any}
 */
proto.v1.Client.prototype.addJwks = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 8, opt_value, proto.google.protobuf.Any, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.v1.Client} returns this
 */
proto.v1.Client.prototype.clearJwksList = function() {
  return this.setJwksList([]);
};


/**
 * optional string jwks_uri = 9;
 * @return {string}
 */
proto.v1.Client.prototype.getJwksUri = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 9, ""));
};


/**
 * @param {string} value
 * @return {!proto.v1.Client} returns this
 */
proto.v1.Client.prototype.setJwksUri = function(value) {
  return jspb.Message.setProto3StringField(this, 9, value);
};


/**
 * optional string logo_uri = 10;
 * @return {string}
 */
proto.v1.Client.prototype.getLogoUri = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 10, ""));
};


/**
 * @param {string} value
 * @return {!proto.v1.Client} returns this
 */
proto.v1.Client.prototype.setLogoUri = function(value) {
  return jspb.Message.setProto3StringField(this, 10, value);
};


/**
 * optional string name = 11;
 * @return {string}
 */
proto.v1.Client.prototype.getName = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 11, ""));
};


/**
 * @param {string} value
 * @return {!proto.v1.Client} returns this
 */
proto.v1.Client.prototype.setName = function(value) {
  return jspb.Message.setProto3StringField(this, 11, value);
};


/**
 * optional string policy_uri = 12;
 * @return {string}
 */
proto.v1.Client.prototype.getPolicyUri = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 12, ""));
};


/**
 * @param {string} value
 * @return {!proto.v1.Client} returns this
 */
proto.v1.Client.prototype.setPolicyUri = function(value) {
  return jspb.Message.setProto3StringField(this, 12, value);
};


/**
 * optional PublicKeysConfiguration public_keys_configuration = 13;
 * @return {!proto.v1.PublicKeysConfiguration}
 */
proto.v1.Client.prototype.getPublicKeysConfiguration = function() {
  return /** @type {!proto.v1.PublicKeysConfiguration} */ (jspb.Message.getFieldWithDefault(this, 13, 0));
};


/**
 * @param {!proto.v1.PublicKeysConfiguration} value
 * @return {!proto.v1.Client} returns this
 */
proto.v1.Client.prototype.setPublicKeysConfiguration = function(value) {
  return jspb.Message.setProto3EnumField(this, 13, value);
};


/**
 * repeated string redirect_uris = 14;
 * @return {!Array<string>}
 */
proto.v1.Client.prototype.getRedirectUrisList = function() {
  return /** @type {!Array<string>} */ (jspb.Message.getRepeatedField(this, 14));
};


/**
 * @param {!Array<string>} value
 * @return {!proto.v1.Client} returns this
 */
proto.v1.Client.prototype.setRedirectUrisList = function(value) {
  return jspb.Message.setField(this, 14, value || []);
};


/**
 * @param {string} value
 * @param {number=} opt_index
 * @return {!proto.v1.Client} returns this
 */
proto.v1.Client.prototype.addRedirectUris = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 14, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.v1.Client} returns this
 */
proto.v1.Client.prototype.clearRedirectUrisList = function() {
  return this.setRedirectUrisList([]);
};


/**
 * repeated RefreshTokenRotationType refresh_token_rotation_type = 15;
 * @return {!Array<!proto.v1.RefreshTokenRotationType>}
 */
proto.v1.Client.prototype.getRefreshTokenRotationTypeList = function() {
  return /** @type {!Array<!proto.v1.RefreshTokenRotationType>} */ (jspb.Message.getRepeatedField(this, 15));
};


/**
 * @param {!Array<!proto.v1.RefreshTokenRotationType>} value
 * @return {!proto.v1.Client} returns this
 */
proto.v1.Client.prototype.setRefreshTokenRotationTypeList = function(value) {
  return jspb.Message.setField(this, 15, value || []);
};


/**
 * @param {!proto.v1.RefreshTokenRotationType} value
 * @param {number=} opt_index
 * @return {!proto.v1.Client} returns this
 */
proto.v1.Client.prototype.addRefreshTokenRotationType = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 15, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.v1.Client} returns this
 */
proto.v1.Client.prototype.clearRefreshTokenRotationTypeList = function() {
  return this.setRefreshTokenRotationTypeList([]);
};


/**
 * repeated ResponseType response_types = 16;
 * @return {!Array<!proto.v1.ResponseType>}
 */
proto.v1.Client.prototype.getResponseTypesList = function() {
  return /** @type {!Array<!proto.v1.ResponseType>} */ (jspb.Message.getRepeatedField(this, 16));
};


/**
 * @param {!Array<!proto.v1.ResponseType>} value
 * @return {!proto.v1.Client} returns this
 */
proto.v1.Client.prototype.setResponseTypesList = function(value) {
  return jspb.Message.setField(this, 16, value || []);
};


/**
 * @param {!proto.v1.ResponseType} value
 * @param {number=} opt_index
 * @return {!proto.v1.Client} returns this
 */
proto.v1.Client.prototype.addResponseTypes = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 16, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.v1.Client} returns this
 */
proto.v1.Client.prototype.clearResponseTypesList = function() {
  return this.setResponseTypesList([]);
};


/**
 * optional string secret = 17;
 * @return {string}
 */
proto.v1.Client.prototype.getSecret = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 17, ""));
};


/**
 * @param {string} value
 * @return {!proto.v1.Client} returns this
 */
proto.v1.Client.prototype.setSecret = function(value) {
  return jspb.Message.setProto3StringField(this, 17, value);
};


/**
 * repeated string scope = 18;
 * @return {!Array<string>}
 */
proto.v1.Client.prototype.getScopeList = function() {
  return /** @type {!Array<string>} */ (jspb.Message.getRepeatedField(this, 18));
};


/**
 * @param {!Array<string>} value
 * @return {!proto.v1.Client} returns this
 */
proto.v1.Client.prototype.setScopeList = function(value) {
  return jspb.Message.setField(this, 18, value || []);
};


/**
 * @param {string} value
 * @param {number=} opt_index
 * @return {!proto.v1.Client} returns this
 */
proto.v1.Client.prototype.addScope = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 18, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.v1.Client} returns this
 */
proto.v1.Client.prototype.clearScopeList = function() {
  return this.setScopeList([]);
};


/**
 * optional string software_id = 19;
 * @return {string}
 */
proto.v1.Client.prototype.getSoftwareId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 19, ""));
};


/**
 * @param {string} value
 * @return {!proto.v1.Client} returns this
 */
proto.v1.Client.prototype.setSoftwareId = function(value) {
  return jspb.Message.setProto3StringField(this, 19, value);
};


/**
 * optional string software_version = 20;
 * @return {string}
 */
proto.v1.Client.prototype.getSoftwareVersion = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 20, ""));
};


/**
 * @param {string} value
 * @return {!proto.v1.Client} returns this
 */
proto.v1.Client.prototype.setSoftwareVersion = function(value) {
  return jspb.Message.setProto3StringField(this, 20, value);
};


/**
 * optional ApplicationType type = 21;
 * @return {!proto.v1.ApplicationType}
 */
proto.v1.Client.prototype.getType = function() {
  return /** @type {!proto.v1.ApplicationType} */ (jspb.Message.getFieldWithDefault(this, 21, 0));
};


/**
 * @param {!proto.v1.ApplicationType} value
 * @return {!proto.v1.Client} returns this
 */
proto.v1.Client.prototype.setType = function(value) {
  return jspb.Message.setProto3EnumField(this, 21, value);
};


/**
 * optional string tos_uri = 22;
 * @return {string}
 */
proto.v1.Client.prototype.getTosUri = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 22, ""));
};


/**
 * @param {string} value
 * @return {!proto.v1.Client} returns this
 */
proto.v1.Client.prototype.setTosUri = function(value) {
  return jspb.Message.setProto3StringField(this, 22, value);
};


/**
 * optional TokenEndpointAuthMethod token_endpoint_auth_method = 23;
 * @return {!proto.v1.TokenEndpointAuthMethod}
 */
proto.v1.Client.prototype.getTokenEndpointAuthMethod = function() {
  return /** @type {!proto.v1.TokenEndpointAuthMethod} */ (jspb.Message.getFieldWithDefault(this, 23, 0));
};


/**
 * @param {!proto.v1.TokenEndpointAuthMethod} value
 * @return {!proto.v1.Client} returns this
 */
proto.v1.Client.prototype.setTokenEndpointAuthMethod = function(value) {
  return jspb.Message.setProto3EnumField(this, 23, value);
};


/**
 * optional string uri = 24;
 * @return {string}
 */
proto.v1.Client.prototype.getUri = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 24, ""));
};


/**
 * @param {string} value
 * @return {!proto.v1.Client} returns this
 */
proto.v1.Client.prototype.setUri = function(value) {
  return jspb.Message.setProto3StringField(this, 24, value);
};


/**
 * optional google.protobuf.Timestamp created_at = 25;
 * @return {?proto.google.protobuf.Timestamp}
 */
proto.v1.Client.prototype.getCreatedAt = function() {
  return /** @type{?proto.google.protobuf.Timestamp} */ (
    jspb.Message.getWrapperField(this, google_protobuf_timestamp_pb.Timestamp, 25));
};


/**
 * @param {?proto.google.protobuf.Timestamp|undefined} value
 * @return {!proto.v1.Client} returns this
*/
proto.v1.Client.prototype.setCreatedAt = function(value) {
  return jspb.Message.setWrapperField(this, 25, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.v1.Client} returns this
 */
proto.v1.Client.prototype.clearCreatedAt = function() {
  return this.setCreatedAt(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.v1.Client.prototype.hasCreatedAt = function() {
  return jspb.Message.getField(this, 25) != null;
};


/**
 * optional google.protobuf.Timestamp updated_at = 26;
 * @return {?proto.google.protobuf.Timestamp}
 */
proto.v1.Client.prototype.getUpdatedAt = function() {
  return /** @type{?proto.google.protobuf.Timestamp} */ (
    jspb.Message.getWrapperField(this, google_protobuf_timestamp_pb.Timestamp, 26));
};


/**
 * @param {?proto.google.protobuf.Timestamp|undefined} value
 * @return {!proto.v1.Client} returns this
*/
proto.v1.Client.prototype.setUpdatedAt = function(value) {
  return jspb.Message.setWrapperField(this, 26, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.v1.Client} returns this
 */
proto.v1.Client.prototype.clearUpdatedAt = function() {
  return this.setUpdatedAt(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.v1.Client.prototype.hasUpdatedAt = function() {
  return jspb.Message.getField(this, 26) != null;
};


goog.object.extend(exports, proto.v1);