# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

import Products_pb2 as Products__pb2
from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2


class ProductsStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.GetProductPrices = channel.unary_unary(
                '/Products/GetProductPrices',
                request_serializer=Products__pb2.Product.SerializeToString,
                response_deserializer=Products__pb2.ProductPrices.FromString,
                )
        self.AddNewPrice = channel.unary_unary(
                '/Products/AddNewPrice',
                request_serializer=Products__pb2.ProductNewPrice.SerializeToString,
                response_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
                )
        self.AddNewProduct = channel.unary_unary(
                '/Products/AddNewProduct',
                request_serializer=Products__pb2.Product.SerializeToString,
                response_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
                )
        self.GetAllProducts = channel.unary_unary(
                '/Products/GetAllProducts',
                request_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
                response_deserializer=Products__pb2.ProductList.FromString,
                )


class ProductsServicer(object):
    """Missing associated documentation comment in .proto file."""

    def GetProductPrices(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def AddNewPrice(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def AddNewProduct(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetAllProducts(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_ProductsServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'GetProductPrices': grpc.unary_unary_rpc_method_handler(
                    servicer.GetProductPrices,
                    request_deserializer=Products__pb2.Product.FromString,
                    response_serializer=Products__pb2.ProductPrices.SerializeToString,
            ),
            'AddNewPrice': grpc.unary_unary_rpc_method_handler(
                    servicer.AddNewPrice,
                    request_deserializer=Products__pb2.ProductNewPrice.FromString,
                    response_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
            ),
            'AddNewProduct': grpc.unary_unary_rpc_method_handler(
                    servicer.AddNewProduct,
                    request_deserializer=Products__pb2.Product.FromString,
                    response_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
            ),
            'GetAllProducts': grpc.unary_unary_rpc_method_handler(
                    servicer.GetAllProducts,
                    request_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
                    response_serializer=Products__pb2.ProductList.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'Products', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class Products(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def GetProductPrices(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/Products/GetProductPrices',
            Products__pb2.Product.SerializeToString,
            Products__pb2.ProductPrices.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def AddNewPrice(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/Products/AddNewPrice',
            Products__pb2.ProductNewPrice.SerializeToString,
            google_dot_protobuf_dot_empty__pb2.Empty.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def AddNewProduct(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/Products/AddNewProduct',
            Products__pb2.Product.SerializeToString,
            google_dot_protobuf_dot_empty__pb2.Empty.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetAllProducts(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/Products/GetAllProducts',
            google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
            Products__pb2.ProductList.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
