part of 'websocket_bloc.dart';

@immutable
class WebsocketApiStartState extends Equatable {
  late final WebSocketChannel channel;
  String message = "";

  @override
  List<Object?> get props => [message];
  WebsocketApiStartState(Uri uri) {
    channel = WebSocketChannel.connect(uri);
  }
}
