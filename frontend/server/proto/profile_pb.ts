// @generated by protoc-gen-es v1.0.0 with parameter "target=ts"
// @generated from file profile.proto (package pb, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3, Timestamp } from "@bufbuild/protobuf";
import { Gender } from "./gender_pb.js";

/**
 * @generated from message pb.Profile
 */
export class Profile extends Message<Profile> {
  /**
   * @generated from field: string id = 1;
   */
  id = "";

  /**
   * @generated from field: string user_id = 2;
   */
  userId = "";

  /**
   * @generated from field: google.protobuf.Timestamp birthdate = 3;
   */
  birthdate?: Timestamp;

  /**
   * @generated from field: pb.Gender gender = 4;
   */
  gender = Gender.UNSPECIFIED;

  /**
   * @generated from field: string locale = 5;
   */
  locale = "";

  /**
   * @generated from field: string given_name = 6;
   */
  givenName = "";

  /**
   * @generated from field: string middle_name = 7;
   */
  middleName = "";

  /**
   * @generated from field: string nickname = 8;
   */
  nickname = "";

  /**
   * @generated from field: string profile = 9;
   */
  profile = "";

  /**
   * @generated from field: string picture = 10;
   */
  picture = "";

  /**
   * @generated from field: string website = 11;
   */
  website = "";

  /**
   * @generated from field: string zoneinfo = 12;
   */
  zoneinfo = "";

  /**
   * @generated from field: google.protobuf.Timestamp created_at = 13;
   */
  createdAt?: Timestamp;

  /**
   * @generated from field: google.protobuf.Timestamp updated_at = 14;
   */
  updatedAt?: Timestamp;

  constructor(data?: PartialMessage<Profile>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "pb.Profile";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "user_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "birthdate", kind: "message", T: Timestamp },
    { no: 4, name: "gender", kind: "enum", T: proto3.getEnumType(Gender) },
    { no: 5, name: "locale", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 6, name: "given_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 7, name: "middle_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 8, name: "nickname", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 9, name: "profile", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 10, name: "picture", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 11, name: "website", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 12, name: "zoneinfo", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 13, name: "created_at", kind: "message", T: Timestamp },
    { no: 14, name: "updated_at", kind: "message", T: Timestamp },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Profile {
    return new Profile().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Profile {
    return new Profile().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Profile {
    return new Profile().fromJsonString(jsonString, options);
  }

  static equals(a: Profile | PlainMessage<Profile> | undefined, b: Profile | PlainMessage<Profile> | undefined): boolean {
    return proto3.util.equals(Profile, a, b);
  }
}

/**
 * @generated from message pb.UpdateProfileRequest
 */
export class UpdateProfileRequest extends Message<UpdateProfileRequest> {
  /**
   * @generated from field: string id = 1;
   */
  id = "";

  /**
   * @generated from field: optional google.protobuf.Timestamp birthdate = 3;
   */
  birthdate?: Timestamp;

  /**
   * @generated from field: optional pb.Gender gender = 4;
   */
  gender?: Gender;

  /**
   * @generated from field: optional string locale = 5;
   */
  locale?: string;

  /**
   * @generated from field: optional string given_name = 6;
   */
  givenName?: string;

  /**
   * @generated from field: optional string middle_name = 7;
   */
  middleName?: string;

  /**
   * @generated from field: optional string nickname = 8;
   */
  nickname?: string;

  /**
   * @generated from field: optional string profile = 9;
   */
  profile?: string;

  /**
   * @generated from field: optional string picture = 10;
   */
  picture?: string;

  /**
   * @generated from field: optional string website = 11;
   */
  website?: string;

  /**
   * @generated from field: optional string zoneinfo = 12;
   */
  zoneinfo?: string;

  constructor(data?: PartialMessage<UpdateProfileRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "pb.UpdateProfileRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "birthdate", kind: "message", T: Timestamp, opt: true },
    { no: 4, name: "gender", kind: "enum", T: proto3.getEnumType(Gender), opt: true },
    { no: 5, name: "locale", kind: "scalar", T: 9 /* ScalarType.STRING */, opt: true },
    { no: 6, name: "given_name", kind: "scalar", T: 9 /* ScalarType.STRING */, opt: true },
    { no: 7, name: "middle_name", kind: "scalar", T: 9 /* ScalarType.STRING */, opt: true },
    { no: 8, name: "nickname", kind: "scalar", T: 9 /* ScalarType.STRING */, opt: true },
    { no: 9, name: "profile", kind: "scalar", T: 9 /* ScalarType.STRING */, opt: true },
    { no: 10, name: "picture", kind: "scalar", T: 9 /* ScalarType.STRING */, opt: true },
    { no: 11, name: "website", kind: "scalar", T: 9 /* ScalarType.STRING */, opt: true },
    { no: 12, name: "zoneinfo", kind: "scalar", T: 9 /* ScalarType.STRING */, opt: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdateProfileRequest {
    return new UpdateProfileRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdateProfileRequest {
    return new UpdateProfileRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdateProfileRequest {
    return new UpdateProfileRequest().fromJsonString(jsonString, options);
  }

  static equals(a: UpdateProfileRequest | PlainMessage<UpdateProfileRequest> | undefined, b: UpdateProfileRequest | PlainMessage<UpdateProfileRequest> | undefined): boolean {
    return proto3.util.equals(UpdateProfileRequest, a, b);
  }
}

