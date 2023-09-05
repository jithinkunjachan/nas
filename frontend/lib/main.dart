import 'dart:convert';

import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:frontend/bloc/api_bloc.dart';
import 'package:frontend/bloc/websocket_bloc.dart';
import 'package:frontend/widgets/textArea.dart';

void main() {
  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  // This widget is the root of your application.
  @override
  Widget build(BuildContext context) {
    return MultiBlocProvider(
      providers: [
        BlocProvider(create: (_) => WebsocketBloc()),
        BlocProvider(create: (_) => ApiBloc())
      ],
      child: MaterialApp(
        title: 'NAS',
        theme: ThemeData(),
        home: const MyHomePage(),
      ),
    );
  }
}

class MyHomePage extends StatefulWidget {
  const MyHomePage({
    super.key,
  });

  @override
  State<MyHomePage> createState() => _MyHomePageState();
}

class _MyHomePageState extends State<MyHomePage> {
  final TextArea textArea = TextArea();
  @override
  void initState() {
    super.initState();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        backgroundColor: Color.fromRGBO(255, 255, 255, 0.493),
      ),
      drawer: Drawer(
          child: ListView(
        children: [
          const ListTile(title: Text("Menus")),
          ListTile(
            title: const Text("lsblk"),
            onTap: () {
              context.read<ApiBloc>().add(SnapRaidSyncEvent("lsblk"));
              Navigator.pop(context);
            },
          ),
          ListTile(
            title: const Text("blkid"),
            onTap: () {
              context.read<ApiBloc>().add(SnapRaidSyncEvent("blkid"));
              Navigator.pop(context);
            },
          ),
          ListTile(
            title: const Text("snapraid status"),
            onTap: () {
              context.read<ApiBloc>().add(SnapRaidSyncEvent("snapraid/status"));
              Navigator.pop(context);
            },
          ),
          ListTile(
            title: const Text("snapraid diff"),
            onTap: () {
              context.read<ApiBloc>().add(SnapRaidSyncEvent("snapraid/diff"));
              Navigator.pop(context);
            },
          ),
          ListTile(
            title: const Text("snapraid sync"),
            onTap: () {
              context.read<ApiBloc>().add(SnapRaidSyncEvent("snapraid/sync"));
              Navigator.pop(context);
            },
          ),
          ListTile(
            title: const Text("snapraid scrub"),
            onTap: () {
              context.read<ApiBloc>().add(SnapRaidSyncEvent("snapraid/scrub"));
              Navigator.pop(context);
            },
          ),
          ListTile(
            title: const Text("snapraid list"),
            onTap: () {
              context.read<ApiBloc>().add(SnapRaidSyncEvent("snapraid/list"));
              Navigator.pop(context);
            },
          ),
          ListTile(
            title: const Text("snapraid dup"),
            onTap: () {
              context.read<ApiBloc>().add(SnapRaidSyncEvent("snapraid/dup"));
              Navigator.pop(context);
            },
          ),
          ListTile(
            title: const Text("snapraid smart"),
            onTap: () {
              context.read<ApiBloc>().add(SnapRaidSyncEvent("snapraid/smart"));
              Navigator.pop(context);
            },
          ),
          ListTile(
            title: const Text("snapraid check"),
            onTap: () {
              context.read<ApiBloc>().add(SnapRaidSyncEvent("snapraid/check"));
              Navigator.pop(context);
            },
          ),
        ],
      )),
      body: Center(
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.stretch,
          children: <Widget>[
            Row(
              mainAxisAlignment: MainAxisAlignment.spaceEvenly,
              children: [
                ElevatedButton(
                  onPressed: () {
                    context
                        .read<ApiBloc>()
                        .add(SnapRaidSyncEvent("snapraid/status"));
                  },
                  child: const Text("status"),
                ),
                ElevatedButton(
                  onPressed: () {
                    context
                        .read<ApiBloc>()
                        .add(SnapRaidSyncEvent("snapraid/diff"));
                  },
                  child: const Text("diff"),
                ),
                ElevatedButton(
                  onPressed: () {
                    context
                        .read<ApiBloc>()
                        .add(SnapRaidSyncEvent("snapraid/sync"));
                  },
                  child: const Text("sync"),
                ),
                ElevatedButton(
                  onPressed: () {
                    context
                        .read<ApiBloc>()
                        .add(SnapRaidSyncEvent("snapraid/scrub"));
                  },
                  child: const Text("scrub"),
                ),
                ElevatedButton(
                  onPressed: () {
                    context
                        .read<ApiBloc>()
                        .add(SnapRaidSyncEvent("snapraid/list"));
                  },
                  child: const Text("list"),
                ),
              ],
            ),
            Row(
              mainAxisAlignment: MainAxisAlignment.spaceEvenly,
              children: [
                ElevatedButton(
                  onPressed: () {
                    context
                        .read<ApiBloc>()
                        .add(SnapRaidSyncEvent("snapraid/dup"));
                  },
                  child: const Text("dup"),
                ),
                ElevatedButton(
                  onPressed: () {
                    context
                        .read<ApiBloc>()
                        .add(SnapRaidSyncEvent("snapraid/smart"));
                  },
                  child: const Text("smart"),
                ),
                ElevatedButton(
                  onPressed: () {
                    context
                        .read<ApiBloc>()
                        .add(SnapRaidSyncEvent("snapraid/check"));
                  },
                  child: const Text("check"),
                ),
              ],
            ),
            BlocBuilder<WebsocketBloc, WebsocketApiStartState>(
                builder: ((context, state) {
              return StreamBuilder(
                  stream: state.channel.stream,
                  builder: ((context, snapshot) {
                    if (snapshot.hasData) {
                      final json = jsonDecode(snapshot.data);
                      final msgType = json["MsgType"];
                      if (msgType == 1) {
                        textArea.Write(json["Message"]);
                      } else if (msgType == 0) {
                        textArea.Clear();
                      }
                    }
                    return textArea;
                  }));
            })),
          ],
        ),
      ),
    );
  }
}
