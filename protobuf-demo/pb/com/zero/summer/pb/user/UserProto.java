// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: proto/user/user.proto

package com.zero.summer.pb.user;

public final class UserProto {
  private UserProto() {}
  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistryLite registry) {
  }

  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistry registry) {
    registerAllExtensions(
        (com.google.protobuf.ExtensionRegistryLite) registry);
  }
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_user_UserRequest_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_user_UserRequest_fieldAccessorTable;
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_user_UserRequest_UserInfoEntry_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_user_UserRequest_UserInfoEntry_fieldAccessorTable;

  public static com.google.protobuf.Descriptors.FileDescriptor
      getDescriptor() {
    return descriptor;
  }
  private static  com.google.protobuf.Descriptors.FileDescriptor
      descriptor;
  static {
    java.lang.String[] descriptorData = {
      "\n\025proto/user/user.proto\022\004user\032\033google/pr" +
      "otobuf/empty.proto\032\037google/protobuf/time" +
      "stamp.proto\"\255\002\n\013UserRequest\022\n\n\002id\030\001 \001(\004\022" +
      "\021\n\tuser_name\030\002 \001(\t\022\020\n\010password\030\003 \001(\t\022\016\n\006" +
      "locked\030\004 \001(\010\022\013\n\003age\030\005 \001(\005\022\013\n\003buf\030\006 \001(\014\022-" +
      "\n\tsend_time\030\007 \001(\0132\032.google.protobuf.Time" +
      "stamp\022\022\n\005token\030\010 \001(\tH\000\210\001\001\022\021\n\tname_list\030\t" +
      " \003(\t\0222\n\tuser_info\030\n \003(\0132\037.user.UserReque" +
      "st.UserInfoEntry\032/\n\rUserInfoEntry\022\013\n\003key" +
      "\030\001 \001(\t\022\r\n\005value\030\002 \001(\022:\0028\001B\010\n\006_token2:\n\004U" +
      "ser\0222\n\005Users\022\021.user.UserRequest\032\026.google" +
      ".protobuf.EmptyB8\n\027com.zero.summer.pb.us" +
      "erB\tUserProtoH\001P\001Z\013./user;user\210\001\001b\006proto" +
      "3"
    };
    descriptor = com.google.protobuf.Descriptors.FileDescriptor
      .internalBuildGeneratedFileFrom(descriptorData,
        new com.google.protobuf.Descriptors.FileDescriptor[] {
          com.google.protobuf.EmptyProto.getDescriptor(),
          com.google.protobuf.TimestampProto.getDescriptor(),
        });
    internal_static_user_UserRequest_descriptor =
      getDescriptor().getMessageTypes().get(0);
    internal_static_user_UserRequest_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_user_UserRequest_descriptor,
        new java.lang.String[] { "Id", "UserName", "Password", "Locked", "Age", "Buf", "SendTime", "Token", "NameList", "UserInfo", "Token", });
    internal_static_user_UserRequest_UserInfoEntry_descriptor =
      internal_static_user_UserRequest_descriptor.getNestedTypes().get(0);
    internal_static_user_UserRequest_UserInfoEntry_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_user_UserRequest_UserInfoEntry_descriptor,
        new java.lang.String[] { "Key", "Value", });
    com.google.protobuf.EmptyProto.getDescriptor();
    com.google.protobuf.TimestampProto.getDescriptor();
  }

  // @@protoc_insertion_point(outer_class_scope)
}
