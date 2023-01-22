# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

import Products_pb2 as Products__pb2
import Users_pb2 as Users__pb2


class UsersStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.GetProducts = channel.unary_unary(
                '/Users/GetProducts',
                request_serializer=Users__pb2.User.SerializeToString,
                response_deserializer=Products__pb2.ProductList.FromString,
                )
        self.AddProduct = channel.unary_unary(
                '/Users/AddProduct',
                request_serializer=Users__pb2.UserProduct.SerializeToString,
                response_deserializer=Products__pb2.ProductList.FromString,
                )


class UsersServicer(object):
    """Missing associated documentation comment in .proto file."""

    def GetProducts(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def AddProduct(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_UsersServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'GetProducts': grpc.unary_unary_rpc_method_handler(
                    servicer.GetProducts,
                    request_deserializer=Users__pb2.User.FromString,
                    response_serializer=Products__pb2.ProductList.SerializeToString,
            ),
            'AddProduct': grpc.unary_unary_rpc_method_handler(
                    servicer.AddProduct,
                    request_deserializer=Users__pb2.UserProduct.FromString,
                    response_serializer=Products__pb2.ProductList.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'Users', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class Users(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def GetProducts(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/Users/GetProducts',
            Users__pb2.User.SerializeToString,
            Products__pb2.ProductList.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def AddProduct(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/Users/AddProduct',
            Users__pb2.UserProduct.SerializeToString,
            Products__pb2.ProductList.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
