// source: drite/account/v1/profile.proto
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
var global =
    (typeof globalThis !== 'undefined' && globalThis) ||
    (typeof window !== 'undefined' && window) ||
    (typeof global !== 'undefined' && global) ||
    (typeof self !== 'undefined' && self) ||
    (function () { return this; }).call(null) ||
    Function('return this')();

var drite_account_v1_gender_pb = require('../../../drite/account/v1/gender_pb.js');
goog.object.extend(proto, drite_account_v1_gender_pb);
var google_protobuf_timestamp_pb = require('google-protobuf/google/protobuf/timestamp_pb.js');
goog.object.extend(proto, google_protobuf_timestamp_pb);
goog.exportSymbol('proto.drite.account.v1.Profile', null, global);
goog.exportSymbol('proto.drite.account.v1.UpdateProfileRequest', null, global);
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
proto.drite.account.v1.Profile = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.drite.account.v1.Profile, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.drite.account.v1.Profile.displayName = 'proto.drite.account.v1.Profile';
}
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
proto.drite.account.v1.UpdateProfileRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.drite.account.v1.UpdateProfileRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.drite.account.v1.UpdateProfileRequest.displayName = 'proto.drite.account.v1.UpdateProfileRequest';
}



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
proto.drite.account.v1.Profile.prototype.toObject = function(opt_includeInstance) {
  return proto.drite.account.v1.Profile.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.drite.account.v1.Profile} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.drite.account.v1.Profile.toObject = function(includeInstance, msg) {
  var f, obj = {
    id: jspb.Message.getFieldWithDefault(msg, 1, ""),
    userId: jspb.Message.getFieldWithDefault(msg, 2, ""),
    birthdate: (f = msg.getBirthdate()) && google_protobuf_timestamp_pb.Timestamp.toObject(includeInstance, f),
    gender: jspb.Message.getFieldWithDefault(msg, 4, 0),
    locale: jspb.Message.getFieldWithDefault(msg, 5, ""),
    givenName: jspb.Message.getFieldWithDefault(msg, 6, ""),
    middleName: jspb.Message.getFieldWithDefault(msg, 7, ""),
    nickname: jspb.Message.getFieldWithDefault(msg, 8, ""),
    profile: jspb.Message.getFieldWithDefault(msg, 9, ""),
    picture: jspb.Message.getFieldWithDefault(msg, 10, ""),
    website: jspb.Message.getFieldWithDefault(msg, 11, ""),
    zoneinfo: jspb.Message.getFieldWithDefault(msg, 12, ""),
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
 * @return {!proto.drite.account.v1.Profile}
 */
proto.drite.account.v1.Profile.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.drite.account.v1.Profile;
  return proto.drite.account.v1.Profile.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.drite.account.v1.Profile} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.drite.account.v1.Profile}
 */
proto.drite.account.v1.Profile.deserializeBinaryFromReader = function(msg, reader) {
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
      var value = new google_protobuf_timestamp_pb.Timestamp;
      reader.readMessage(value,google_protobuf_timestamp_pb.Timestamp.deserializeBinaryFromReader);
      msg.setBirthdate(value);
      break;
    case 4:
      var value = /** @type {!proto.drite.account.v1.Gender} */ (reader.readEnum());
      msg.setGender(value);
      break;
    case 5:
      var value = /** @type {string} */ (reader.readString());
      msg.setLocale(value);
      break;
    case 6:
      var value = /** @type {string} */ (reader.readString());
      msg.setGivenName(value);
      break;
    case 7:
      var value = /** @type {string} */ (reader.readString());
      msg.setMiddleName(value);
      break;
    case 8:
      var value = /** @type {string} */ (reader.readString());
      msg.setNickname(value);
      break;
    case 9:
      var value = /** @type {string} */ (reader.readString());
      msg.setProfile(value);
      break;
    case 10:
      var value = /** @type {string} */ (reader.readString());
      msg.setPicture(value);
      break;
    case 11:
      var value = /** @type {string} */ (reader.readString());
      msg.setWebsite(value);
      break;
    case 12:
      var value = /** @type {string} */ (reader.readString());
      msg.setZoneinfo(value);
      break;
    case 13:
      var value = new google_protobuf_timestamp_pb.Timestamp;
      reader.readMessage(value,google_protobuf_timestamp_pb.Timestamp.deserializeBinaryFromReader);
      msg.setCreatedAt(value);
      break;
    case 14:
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
proto.drite.account.v1.Profile.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.drite.account.v1.Profile.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.drite.account.v1.Profile} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.drite.account.v1.Profile.serializeBinaryToWriter = function(message, writer) {
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
  f = message.getBirthdate();
  if (f != null) {
    writer.writeMessage(
      3,
      f,
      google_protobuf_timestamp_pb.Timestamp.serializeBinaryToWriter
    );
  }
  f = message.getGender();
  if (f !== 0.0) {
    writer.writeEnum(
      4,
      f
    );
  }
  f = message.getLocale();
  if (f.length > 0) {
    writer.writeString(
      5,
      f
    );
  }
  f = message.getGivenName();
  if (f.length > 0) {
    writer.writeString(
      6,
      f
    );
  }
  f = message.getMiddleName();
  if (f.length > 0) {
    writer.writeString(
      7,
      f
    );
  }
  f = message.getNickname();
  if (f.length > 0) {
    writer.writeString(
      8,
      f
    );
  }
  f = message.getProfile();
  if (f.length > 0) {
    writer.writeString(
      9,
      f
    );
  }
  f = message.getPicture();
  if (f.length > 0) {
    writer.writeString(
      10,
      f
    );
  }
  f = message.getWebsite();
  if (f.length > 0) {
    writer.writeString(
      11,
      f
    );
  }
  f = message.getZoneinfo();
  if (f.length > 0) {
    writer.writeString(
      12,
      f
    );
  }
  f = message.getCreatedAt();
  if (f != null) {
    writer.writeMessage(
      13,
      f,
      google_protobuf_timestamp_pb.Timestamp.serializeBinaryToWriter
    );
  }
  f = message.getUpdatedAt();
  if (f != null) {
    writer.writeMessage(
      14,
      f,
      google_protobuf_timestamp_pb.Timestamp.serializeBinaryToWriter
    );
  }
};


/**
 * optional string id = 1;
 * @return {string}
 */
proto.drite.account.v1.Profile.prototype.getId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.drite.account.v1.Profile} returns this
 */
proto.drite.account.v1.Profile.prototype.setId = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string user_id = 2;
 * @return {string}
 */
proto.drite.account.v1.Profile.prototype.getUserId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.drite.account.v1.Profile} returns this
 */
proto.drite.account.v1.Profile.prototype.setUserId = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional google.protobuf.Timestamp birthdate = 3;
 * @return {?proto.google.protobuf.Timestamp}
 */
proto.drite.account.v1.Profile.prototype.getBirthdate = function() {
  return /** @type{?proto.google.protobuf.Timestamp} */ (
    jspb.Message.getWrapperField(this, google_protobuf_timestamp_pb.Timestamp, 3));
};


/**
 * @param {?proto.google.protobuf.Timestamp|undefined} value
 * @return {!proto.drite.account.v1.Profile} returns this
*/
proto.drite.account.v1.Profile.prototype.setBirthdate = function(value) {
  return jspb.Message.setWrapperField(this, 3, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.drite.account.v1.Profile} returns this
 */
proto.drite.account.v1.Profile.prototype.clearBirthdate = function() {
  return this.setBirthdate(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.drite.account.v1.Profile.prototype.hasBirthdate = function() {
  return jspb.Message.getField(this, 3) != null;
};


/**
 * optional Gender gender = 4;
 * @return {!proto.drite.account.v1.Gender}
 */
proto.drite.account.v1.Profile.prototype.getGender = function() {
  return /** @type {!proto.drite.account.v1.Gender} */ (jspb.Message.getFieldWithDefault(this, 4, 0));
};


/**
 * @param {!proto.drite.account.v1.Gender} value
 * @return {!proto.drite.account.v1.Profile} returns this
 */
proto.drite.account.v1.Profile.prototype.setGender = function(value) {
  return jspb.Message.setProto3EnumField(this, 4, value);
};


/**
 * optional string locale = 5;
 * @return {string}
 */
proto.drite.account.v1.Profile.prototype.getLocale = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 5, ""));
};


/**
 * @param {string} value
 * @return {!proto.drite.account.v1.Profile} returns this
 */
proto.drite.account.v1.Profile.prototype.setLocale = function(value) {
  return jspb.Message.setProto3StringField(this, 5, value);
};


/**
 * optional string given_name = 6;
 * @return {string}
 */
proto.drite.account.v1.Profile.prototype.getGivenName = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 6, ""));
};


/**
 * @param {string} value
 * @return {!proto.drite.account.v1.Profile} returns this
 */
proto.drite.account.v1.Profile.prototype.setGivenName = function(value) {
  return jspb.Message.setProto3StringField(this, 6, value);
};


/**
 * optional string middle_name = 7;
 * @return {string}
 */
proto.drite.account.v1.Profile.prototype.getMiddleName = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 7, ""));
};


/**
 * @param {string} value
 * @return {!proto.drite.account.v1.Profile} returns this
 */
proto.drite.account.v1.Profile.prototype.setMiddleName = function(value) {
  return jspb.Message.setProto3StringField(this, 7, value);
};


/**
 * optional string nickname = 8;
 * @return {string}
 */
proto.drite.account.v1.Profile.prototype.getNickname = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 8, ""));
};


/**
 * @param {string} value
 * @return {!proto.drite.account.v1.Profile} returns this
 */
proto.drite.account.v1.Profile.prototype.setNickname = function(value) {
  return jspb.Message.setProto3StringField(this, 8, value);
};


/**
 * optional string profile = 9;
 * @return {string}
 */
proto.drite.account.v1.Profile.prototype.getProfile = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 9, ""));
};


/**
 * @param {string} value
 * @return {!proto.drite.account.v1.Profile} returns this
 */
proto.drite.account.v1.Profile.prototype.setProfile = function(value) {
  return jspb.Message.setProto3StringField(this, 9, value);
};


/**
 * optional string picture = 10;
 * @return {string}
 */
proto.drite.account.v1.Profile.prototype.getPicture = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 10, ""));
};


/**
 * @param {string} value
 * @return {!proto.drite.account.v1.Profile} returns this
 */
proto.drite.account.v1.Profile.prototype.setPicture = function(value) {
  return jspb.Message.setProto3StringField(this, 10, value);
};


/**
 * optional string website = 11;
 * @return {string}
 */
proto.drite.account.v1.Profile.prototype.getWebsite = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 11, ""));
};


/**
 * @param {string} value
 * @return {!proto.drite.account.v1.Profile} returns this
 */
proto.drite.account.v1.Profile.prototype.setWebsite = function(value) {
  return jspb.Message.setProto3StringField(this, 11, value);
};


/**
 * optional string zoneinfo = 12;
 * @return {string}
 */
proto.drite.account.v1.Profile.prototype.getZoneinfo = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 12, ""));
};


/**
 * @param {string} value
 * @return {!proto.drite.account.v1.Profile} returns this
 */
proto.drite.account.v1.Profile.prototype.setZoneinfo = function(value) {
  return jspb.Message.setProto3StringField(this, 12, value);
};


/**
 * optional google.protobuf.Timestamp created_at = 13;
 * @return {?proto.google.protobuf.Timestamp}
 */
proto.drite.account.v1.Profile.prototype.getCreatedAt = function() {
  return /** @type{?proto.google.protobuf.Timestamp} */ (
    jspb.Message.getWrapperField(this, google_protobuf_timestamp_pb.Timestamp, 13));
};


/**
 * @param {?proto.google.protobuf.Timestamp|undefined} value
 * @return {!proto.drite.account.v1.Profile} returns this
*/
proto.drite.account.v1.Profile.prototype.setCreatedAt = function(value) {
  return jspb.Message.setWrapperField(this, 13, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.drite.account.v1.Profile} returns this
 */
proto.drite.account.v1.Profile.prototype.clearCreatedAt = function() {
  return this.setCreatedAt(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.drite.account.v1.Profile.prototype.hasCreatedAt = function() {
  return jspb.Message.getField(this, 13) != null;
};


/**
 * optional google.protobuf.Timestamp updated_at = 14;
 * @return {?proto.google.protobuf.Timestamp}
 */
proto.drite.account.v1.Profile.prototype.getUpdatedAt = function() {
  return /** @type{?proto.google.protobuf.Timestamp} */ (
    jspb.Message.getWrapperField(this, google_protobuf_timestamp_pb.Timestamp, 14));
};


/**
 * @param {?proto.google.protobuf.Timestamp|undefined} value
 * @return {!proto.drite.account.v1.Profile} returns this
*/
proto.drite.account.v1.Profile.prototype.setUpdatedAt = function(value) {
  return jspb.Message.setWrapperField(this, 14, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.drite.account.v1.Profile} returns this
 */
proto.drite.account.v1.Profile.prototype.clearUpdatedAt = function() {
  return this.setUpdatedAt(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.drite.account.v1.Profile.prototype.hasUpdatedAt = function() {
  return jspb.Message.getField(this, 14) != null;
};





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
proto.drite.account.v1.UpdateProfileRequest.prototype.toObject = function(opt_includeInstance) {
  return proto.drite.account.v1.UpdateProfileRequest.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.drite.account.v1.UpdateProfileRequest} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.drite.account.v1.UpdateProfileRequest.toObject = function(includeInstance, msg) {
  var f, obj = {
    id: jspb.Message.getFieldWithDefault(msg, 1, ""),
    birthdate: (f = msg.getBirthdate()) && google_protobuf_timestamp_pb.Timestamp.toObject(includeInstance, f),
    gender: jspb.Message.getFieldWithDefault(msg, 4, 0),
    locale: jspb.Message.getFieldWithDefault(msg, 5, ""),
    givenName: jspb.Message.getFieldWithDefault(msg, 6, ""),
    middleName: jspb.Message.getFieldWithDefault(msg, 7, ""),
    nickname: jspb.Message.getFieldWithDefault(msg, 8, ""),
    profile: jspb.Message.getFieldWithDefault(msg, 9, ""),
    picture: jspb.Message.getFieldWithDefault(msg, 10, ""),
    website: jspb.Message.getFieldWithDefault(msg, 11, ""),
    zoneinfo: jspb.Message.getFieldWithDefault(msg, 12, "")
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
 * @return {!proto.drite.account.v1.UpdateProfileRequest}
 */
proto.drite.account.v1.UpdateProfileRequest.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.drite.account.v1.UpdateProfileRequest;
  return proto.drite.account.v1.UpdateProfileRequest.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.drite.account.v1.UpdateProfileRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.drite.account.v1.UpdateProfileRequest}
 */
proto.drite.account.v1.UpdateProfileRequest.deserializeBinaryFromReader = function(msg, reader) {
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
    case 3:
      var value = new google_protobuf_timestamp_pb.Timestamp;
      reader.readMessage(value,google_protobuf_timestamp_pb.Timestamp.deserializeBinaryFromReader);
      msg.setBirthdate(value);
      break;
    case 4:
      var value = /** @type {!proto.drite.account.v1.Gender} */ (reader.readEnum());
      msg.setGender(value);
      break;
    case 5:
      var value = /** @type {string} */ (reader.readString());
      msg.setLocale(value);
      break;
    case 6:
      var value = /** @type {string} */ (reader.readString());
      msg.setGivenName(value);
      break;
    case 7:
      var value = /** @type {string} */ (reader.readString());
      msg.setMiddleName(value);
      break;
    case 8:
      var value = /** @type {string} */ (reader.readString());
      msg.setNickname(value);
      break;
    case 9:
      var value = /** @type {string} */ (reader.readString());
      msg.setProfile(value);
      break;
    case 10:
      var value = /** @type {string} */ (reader.readString());
      msg.setPicture(value);
      break;
    case 11:
      var value = /** @type {string} */ (reader.readString());
      msg.setWebsite(value);
      break;
    case 12:
      var value = /** @type {string} */ (reader.readString());
      msg.setZoneinfo(value);
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
proto.drite.account.v1.UpdateProfileRequest.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.drite.account.v1.UpdateProfileRequest.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.drite.account.v1.UpdateProfileRequest} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.drite.account.v1.UpdateProfileRequest.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getId();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getBirthdate();
  if (f != null) {
    writer.writeMessage(
      3,
      f,
      google_protobuf_timestamp_pb.Timestamp.serializeBinaryToWriter
    );
  }
  f = /** @type {!proto.drite.account.v1.Gender} */ (jspb.Message.getField(message, 4));
  if (f != null) {
    writer.writeEnum(
      4,
      f
    );
  }
  f = /** @type {string} */ (jspb.Message.getField(message, 5));
  if (f != null) {
    writer.writeString(
      5,
      f
    );
  }
  f = /** @type {string} */ (jspb.Message.getField(message, 6));
  if (f != null) {
    writer.writeString(
      6,
      f
    );
  }
  f = /** @type {string} */ (jspb.Message.getField(message, 7));
  if (f != null) {
    writer.writeString(
      7,
      f
    );
  }
  f = /** @type {string} */ (jspb.Message.getField(message, 8));
  if (f != null) {
    writer.writeString(
      8,
      f
    );
  }
  f = /** @type {string} */ (jspb.Message.getField(message, 9));
  if (f != null) {
    writer.writeString(
      9,
      f
    );
  }
  f = /** @type {string} */ (jspb.Message.getField(message, 10));
  if (f != null) {
    writer.writeString(
      10,
      f
    );
  }
  f = /** @type {string} */ (jspb.Message.getField(message, 11));
  if (f != null) {
    writer.writeString(
      11,
      f
    );
  }
  f = /** @type {string} */ (jspb.Message.getField(message, 12));
  if (f != null) {
    writer.writeString(
      12,
      f
    );
  }
};


/**
 * optional string id = 1;
 * @return {string}
 */
proto.drite.account.v1.UpdateProfileRequest.prototype.getId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.drite.account.v1.UpdateProfileRequest} returns this
 */
proto.drite.account.v1.UpdateProfileRequest.prototype.setId = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional google.protobuf.Timestamp birthdate = 3;
 * @return {?proto.google.protobuf.Timestamp}
 */
proto.drite.account.v1.UpdateProfileRequest.prototype.getBirthdate = function() {
  return /** @type{?proto.google.protobuf.Timestamp} */ (
    jspb.Message.getWrapperField(this, google_protobuf_timestamp_pb.Timestamp, 3));
};


/**
 * @param {?proto.google.protobuf.Timestamp|undefined} value
 * @return {!proto.drite.account.v1.UpdateProfileRequest} returns this
*/
proto.drite.account.v1.UpdateProfileRequest.prototype.setBirthdate = function(value) {
  return jspb.Message.setWrapperField(this, 3, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.drite.account.v1.UpdateProfileRequest} returns this
 */
proto.drite.account.v1.UpdateProfileRequest.prototype.clearBirthdate = function() {
  return this.setBirthdate(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.drite.account.v1.UpdateProfileRequest.prototype.hasBirthdate = function() {
  return jspb.Message.getField(this, 3) != null;
};


/**
 * optional Gender gender = 4;
 * @return {!proto.drite.account.v1.Gender}
 */
proto.drite.account.v1.UpdateProfileRequest.prototype.getGender = function() {
  return /** @type {!proto.drite.account.v1.Gender} */ (jspb.Message.getFieldWithDefault(this, 4, 0));
};


/**
 * @param {!proto.drite.account.v1.Gender} value
 * @return {!proto.drite.account.v1.UpdateProfileRequest} returns this
 */
proto.drite.account.v1.UpdateProfileRequest.prototype.setGender = function(value) {
  return jspb.Message.setField(this, 4, value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.drite.account.v1.UpdateProfileRequest} returns this
 */
proto.drite.account.v1.UpdateProfileRequest.prototype.clearGender = function() {
  return jspb.Message.setField(this, 4, undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.drite.account.v1.UpdateProfileRequest.prototype.hasGender = function() {
  return jspb.Message.getField(this, 4) != null;
};


/**
 * optional string locale = 5;
 * @return {string}
 */
proto.drite.account.v1.UpdateProfileRequest.prototype.getLocale = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 5, ""));
};


/**
 * @param {string} value
 * @return {!proto.drite.account.v1.UpdateProfileRequest} returns this
 */
proto.drite.account.v1.UpdateProfileRequest.prototype.setLocale = function(value) {
  return jspb.Message.setField(this, 5, value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.drite.account.v1.UpdateProfileRequest} returns this
 */
proto.drite.account.v1.UpdateProfileRequest.prototype.clearLocale = function() {
  return jspb.Message.setField(this, 5, undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.drite.account.v1.UpdateProfileRequest.prototype.hasLocale = function() {
  return jspb.Message.getField(this, 5) != null;
};


/**
 * optional string given_name = 6;
 * @return {string}
 */
proto.drite.account.v1.UpdateProfileRequest.prototype.getGivenName = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 6, ""));
};


/**
 * @param {string} value
 * @return {!proto.drite.account.v1.UpdateProfileRequest} returns this
 */
proto.drite.account.v1.UpdateProfileRequest.prototype.setGivenName = function(value) {
  return jspb.Message.setField(this, 6, value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.drite.account.v1.UpdateProfileRequest} returns this
 */
proto.drite.account.v1.UpdateProfileRequest.prototype.clearGivenName = function() {
  return jspb.Message.setField(this, 6, undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.drite.account.v1.UpdateProfileRequest.prototype.hasGivenName = function() {
  return jspb.Message.getField(this, 6) != null;
};


/**
 * optional string middle_name = 7;
 * @return {string}
 */
proto.drite.account.v1.UpdateProfileRequest.prototype.getMiddleName = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 7, ""));
};


/**
 * @param {string} value
 * @return {!proto.drite.account.v1.UpdateProfileRequest} returns this
 */
proto.drite.account.v1.UpdateProfileRequest.prototype.setMiddleName = function(value) {
  return jspb.Message.setField(this, 7, value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.drite.account.v1.UpdateProfileRequest} returns this
 */
proto.drite.account.v1.UpdateProfileRequest.prototype.clearMiddleName = function() {
  return jspb.Message.setField(this, 7, undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.drite.account.v1.UpdateProfileRequest.prototype.hasMiddleName = function() {
  return jspb.Message.getField(this, 7) != null;
};


/**
 * optional string nickname = 8;
 * @return {string}
 */
proto.drite.account.v1.UpdateProfileRequest.prototype.getNickname = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 8, ""));
};


/**
 * @param {string} value
 * @return {!proto.drite.account.v1.UpdateProfileRequest} returns this
 */
proto.drite.account.v1.UpdateProfileRequest.prototype.setNickname = function(value) {
  return jspb.Message.setField(this, 8, value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.drite.account.v1.UpdateProfileRequest} returns this
 */
proto.drite.account.v1.UpdateProfileRequest.prototype.clearNickname = function() {
  return jspb.Message.setField(this, 8, undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.drite.account.v1.UpdateProfileRequest.prototype.hasNickname = function() {
  return jspb.Message.getField(this, 8) != null;
};


/**
 * optional string profile = 9;
 * @return {string}
 */
proto.drite.account.v1.UpdateProfileRequest.prototype.getProfile = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 9, ""));
};


/**
 * @param {string} value
 * @return {!proto.drite.account.v1.UpdateProfileRequest} returns this
 */
proto.drite.account.v1.UpdateProfileRequest.prototype.setProfile = function(value) {
  return jspb.Message.setField(this, 9, value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.drite.account.v1.UpdateProfileRequest} returns this
 */
proto.drite.account.v1.UpdateProfileRequest.prototype.clearProfile = function() {
  return jspb.Message.setField(this, 9, undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.drite.account.v1.UpdateProfileRequest.prototype.hasProfile = function() {
  return jspb.Message.getField(this, 9) != null;
};


/**
 * optional string picture = 10;
 * @return {string}
 */
proto.drite.account.v1.UpdateProfileRequest.prototype.getPicture = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 10, ""));
};


/**
 * @param {string} value
 * @return {!proto.drite.account.v1.UpdateProfileRequest} returns this
 */
proto.drite.account.v1.UpdateProfileRequest.prototype.setPicture = function(value) {
  return jspb.Message.setField(this, 10, value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.drite.account.v1.UpdateProfileRequest} returns this
 */
proto.drite.account.v1.UpdateProfileRequest.prototype.clearPicture = function() {
  return jspb.Message.setField(this, 10, undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.drite.account.v1.UpdateProfileRequest.prototype.hasPicture = function() {
  return jspb.Message.getField(this, 10) != null;
};


/**
 * optional string website = 11;
 * @return {string}
 */
proto.drite.account.v1.UpdateProfileRequest.prototype.getWebsite = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 11, ""));
};


/**
 * @param {string} value
 * @return {!proto.drite.account.v1.UpdateProfileRequest} returns this
 */
proto.drite.account.v1.UpdateProfileRequest.prototype.setWebsite = function(value) {
  return jspb.Message.setField(this, 11, value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.drite.account.v1.UpdateProfileRequest} returns this
 */
proto.drite.account.v1.UpdateProfileRequest.prototype.clearWebsite = function() {
  return jspb.Message.setField(this, 11, undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.drite.account.v1.UpdateProfileRequest.prototype.hasWebsite = function() {
  return jspb.Message.getField(this, 11) != null;
};


/**
 * optional string zoneinfo = 12;
 * @return {string}
 */
proto.drite.account.v1.UpdateProfileRequest.prototype.getZoneinfo = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 12, ""));
};


/**
 * @param {string} value
 * @return {!proto.drite.account.v1.UpdateProfileRequest} returns this
 */
proto.drite.account.v1.UpdateProfileRequest.prototype.setZoneinfo = function(value) {
  return jspb.Message.setField(this, 12, value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.drite.account.v1.UpdateProfileRequest} returns this
 */
proto.drite.account.v1.UpdateProfileRequest.prototype.clearZoneinfo = function() {
  return jspb.Message.setField(this, 12, undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.drite.account.v1.UpdateProfileRequest.prototype.hasZoneinfo = function() {
  return jspb.Message.getField(this, 12) != null;
};


goog.object.extend(exports, proto.drite.account.v1);
