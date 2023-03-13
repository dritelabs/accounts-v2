import { User as UserMessage } from "~~/server/contexts/shared/infrastructure/proto/drite/account/v1/user_pb";
import { UserSerializer } from "../../domain/serializers/user-serializer";

export function defineGRPCUserSerializer(): UserSerializer<UserMessage> {
  return {
    serializeToEntity(input) {
      return {
        id: input.getId(),
        profile: {},
      };
    },
  };
}