part of 'api_bloc.dart';

@immutable
abstract class ApiEvent {}

class SnapRaidSyncEvent extends ApiEvent {
  final String path;
  SnapRaidSyncEvent(this.path);
}
