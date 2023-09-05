part of 'api_bloc.dart';

@immutable
abstract class ApiBlocState {}

class SnapRaidSyncState extends ApiBlocState {
  void call(String path) async {
    final client = http.Client();
    try {
      var response = client.get(Uri.http("192.168.50.103:8081", path));
    } finally {
      client.close();
    }
  }
}

class ApiLoadingState extends ApiBlocState {}
