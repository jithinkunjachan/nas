import 'dart:io';

import 'package:bloc/bloc.dart';
import 'package:meta/meta.dart';
import 'package:http/http.dart' as http;

part 'api_event.dart';
part 'api_state.dart';

class ApiBloc extends Bloc<ApiEvent, ApiBlocState> {
  ApiBloc() : super(SnapRaidSyncState()) {
    on<SnapRaidSyncEvent>((event, emit) {
      final state = SnapRaidSyncState();
      emit(ApiLoadingState());
      state.call(event.path);
      emit(state);
    });
  }
}
