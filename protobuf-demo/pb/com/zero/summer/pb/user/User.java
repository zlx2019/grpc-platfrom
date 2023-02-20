// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: proto/user/user.proto

package com.zero.summer.pb.user;

/**
 * Protobuf service {@code user.User}
 */
public  abstract class User
    implements com.google.protobuf.Service {
  protected User() {}

  public interface Interface {
    /**
     * <code>rpc Users(.user.UserRequest) returns (.google.protobuf.Empty);</code>
     */
    public abstract void users(
        com.google.protobuf.RpcController controller,
        com.zero.summer.pb.user.UserRequest request,
        com.google.protobuf.RpcCallback<com.google.protobuf.Empty> done);

  }

  public static com.google.protobuf.Service newReflectiveService(
      final Interface impl) {
    return new User() {
      @java.lang.Override
      public  void users(
          com.google.protobuf.RpcController controller,
          com.zero.summer.pb.user.UserRequest request,
          com.google.protobuf.RpcCallback<com.google.protobuf.Empty> done) {
        impl.users(controller, request, done);
      }

    };
  }

  public static com.google.protobuf.BlockingService
      newReflectiveBlockingService(final BlockingInterface impl) {
    return new com.google.protobuf.BlockingService() {
      public final com.google.protobuf.Descriptors.ServiceDescriptor
          getDescriptorForType() {
        return getDescriptor();
      }

      public final com.google.protobuf.Message callBlockingMethod(
          com.google.protobuf.Descriptors.MethodDescriptor method,
          com.google.protobuf.RpcController controller,
          com.google.protobuf.Message request)
          throws com.google.protobuf.ServiceException {
        if (method.getService() != getDescriptor()) {
          throw new java.lang.IllegalArgumentException(
            "Service.callBlockingMethod() given method descriptor for " +
            "wrong service type.");
        }
        switch(method.getIndex()) {
          case 0:
            return impl.users(controller, (com.zero.summer.pb.user.UserRequest)request);
          default:
            throw new java.lang.AssertionError("Can't get here.");
        }
      }

      public final com.google.protobuf.Message
          getRequestPrototype(
          com.google.protobuf.Descriptors.MethodDescriptor method) {
        if (method.getService() != getDescriptor()) {
          throw new java.lang.IllegalArgumentException(
            "Service.getRequestPrototype() given method " +
            "descriptor for wrong service type.");
        }
        switch(method.getIndex()) {
          case 0:
            return com.zero.summer.pb.user.UserRequest.getDefaultInstance();
          default:
            throw new java.lang.AssertionError("Can't get here.");
        }
      }

      public final com.google.protobuf.Message
          getResponsePrototype(
          com.google.protobuf.Descriptors.MethodDescriptor method) {
        if (method.getService() != getDescriptor()) {
          throw new java.lang.IllegalArgumentException(
            "Service.getResponsePrototype() given method " +
            "descriptor for wrong service type.");
        }
        switch(method.getIndex()) {
          case 0:
            return com.google.protobuf.Empty.getDefaultInstance();
          default:
            throw new java.lang.AssertionError("Can't get here.");
        }
      }

    };
  }

  /**
   * <code>rpc Users(.user.UserRequest) returns (.google.protobuf.Empty);</code>
   */
  public abstract void users(
      com.google.protobuf.RpcController controller,
      com.zero.summer.pb.user.UserRequest request,
      com.google.protobuf.RpcCallback<com.google.protobuf.Empty> done);

  public static final
      com.google.protobuf.Descriptors.ServiceDescriptor
      getDescriptor() {
    return com.zero.summer.pb.user.UserProto.getDescriptor().getServices().get(0);
  }
  public final com.google.protobuf.Descriptors.ServiceDescriptor
      getDescriptorForType() {
    return getDescriptor();
  }

  public final void callMethod(
      com.google.protobuf.Descriptors.MethodDescriptor method,
      com.google.protobuf.RpcController controller,
      com.google.protobuf.Message request,
      com.google.protobuf.RpcCallback<
        com.google.protobuf.Message> done) {
    if (method.getService() != getDescriptor()) {
      throw new java.lang.IllegalArgumentException(
        "Service.callMethod() given method descriptor for wrong " +
        "service type.");
    }
    switch(method.getIndex()) {
      case 0:
        this.users(controller, (com.zero.summer.pb.user.UserRequest)request,
          com.google.protobuf.RpcUtil.<com.google.protobuf.Empty>specializeCallback(
            done));
        return;
      default:
        throw new java.lang.AssertionError("Can't get here.");
    }
  }

  public final com.google.protobuf.Message
      getRequestPrototype(
      com.google.protobuf.Descriptors.MethodDescriptor method) {
    if (method.getService() != getDescriptor()) {
      throw new java.lang.IllegalArgumentException(
        "Service.getRequestPrototype() given method " +
        "descriptor for wrong service type.");
    }
    switch(method.getIndex()) {
      case 0:
        return com.zero.summer.pb.user.UserRequest.getDefaultInstance();
      default:
        throw new java.lang.AssertionError("Can't get here.");
    }
  }

  public final com.google.protobuf.Message
      getResponsePrototype(
      com.google.protobuf.Descriptors.MethodDescriptor method) {
    if (method.getService() != getDescriptor()) {
      throw new java.lang.IllegalArgumentException(
        "Service.getResponsePrototype() given method " +
        "descriptor for wrong service type.");
    }
    switch(method.getIndex()) {
      case 0:
        return com.google.protobuf.Empty.getDefaultInstance();
      default:
        throw new java.lang.AssertionError("Can't get here.");
    }
  }

  public static Stub newStub(
      com.google.protobuf.RpcChannel channel) {
    return new Stub(channel);
  }

  public static final class Stub extends com.zero.summer.pb.user.User implements Interface {
    private Stub(com.google.protobuf.RpcChannel channel) {
      this.channel = channel;
    }

    private final com.google.protobuf.RpcChannel channel;

    public com.google.protobuf.RpcChannel getChannel() {
      return channel;
    }

    public  void users(
        com.google.protobuf.RpcController controller,
        com.zero.summer.pb.user.UserRequest request,
        com.google.protobuf.RpcCallback<com.google.protobuf.Empty> done) {
      channel.callMethod(
        getDescriptor().getMethods().get(0),
        controller,
        request,
        com.google.protobuf.Empty.getDefaultInstance(),
        com.google.protobuf.RpcUtil.generalizeCallback(
          done,
          com.google.protobuf.Empty.class,
          com.google.protobuf.Empty.getDefaultInstance()));
    }
  }

  public static BlockingInterface newBlockingStub(
      com.google.protobuf.BlockingRpcChannel channel) {
    return new BlockingStub(channel);
  }

  public interface BlockingInterface {
    public com.google.protobuf.Empty users(
        com.google.protobuf.RpcController controller,
        com.zero.summer.pb.user.UserRequest request)
        throws com.google.protobuf.ServiceException;
  }

  private static final class BlockingStub implements BlockingInterface {
    private BlockingStub(com.google.protobuf.BlockingRpcChannel channel) {
      this.channel = channel;
    }

    private final com.google.protobuf.BlockingRpcChannel channel;

    public com.google.protobuf.Empty users(
        com.google.protobuf.RpcController controller,
        com.zero.summer.pb.user.UserRequest request)
        throws com.google.protobuf.ServiceException {
      return (com.google.protobuf.Empty) channel.callBlockingMethod(
        getDescriptor().getMethods().get(0),
        controller,
        request,
        com.google.protobuf.Empty.getDefaultInstance());
    }

  }

  // @@protoc_insertion_point(class_scope:user.User)
}
