// @generated by protoc-gen-es v1.0.0 with parameter "target=ts"
// @generated from file drite/account/v1/refresh_token_rotation_type.proto (package drite.account.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { proto3 } from "@bufbuild/protobuf";

/**
 * @generated from enum drite.account.v1.RefreshTokenRotationType
 */
export enum RefreshTokenRotationType {
  /**
   * @generated from enum value: REFRESH_TOKEN_ROTATION_TYPE_UNSPECIFIED = 0;
   */
  UNSPECIFIED = 0,

  /**
   * @generated from enum value: REFRESH_TOKEN_ROTATION_TYPE_ROTATE = 1;
   */
  ROTATE = 1,

  /**
   * @generated from enum value: REFRESH_TOKEN_ROTATION_TYPE_STATIC = 2;
   */
  STATIC = 2,
}
// Retrieve enum metadata with: proto3.getEnumType(RefreshTokenRotationType)
proto3.util.setEnumType(RefreshTokenRotationType, "drite.account.v1.RefreshTokenRotationType", [
  { no: 0, name: "REFRESH_TOKEN_ROTATION_TYPE_UNSPECIFIED" },
  { no: 1, name: "REFRESH_TOKEN_ROTATION_TYPE_ROTATE" },
  { no: 2, name: "REFRESH_TOKEN_ROTATION_TYPE_STATIC" },
]);

