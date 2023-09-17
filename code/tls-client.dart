class GrpcClient {
  Future<void> init() async {
    var trustedRoot = await rootBundle.load('assets/ca/alfie-cloud-services_com.crt');
    final channel = ClientChannel(
      grpcHost,
      port: grpcPort,
      options: ChannelOptions(credentials: ChannelCredentials.secure(certificates: trustedRoot.buffer.asUint8List())),
    );
    _client = AlfieClient(channel);
  }
}