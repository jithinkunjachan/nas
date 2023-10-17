import 'dart:convert';

import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:flutter_svg/flutter_svg.dart';
import 'package:frontend/bloc/api_bloc.dart';
import 'package:frontend/bloc/websocket_bloc.dart';
import 'package:frontend/widgets/button.dart';
import 'package:frontend/widgets/loglist.dart';

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
  final LogList loglist = LogList();
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
            Padding(
              padding: const EdgeInsets.all(30),
              child: Row(
                mainAxisAlignment: MainAxisAlignment.spaceBetween,
                children: [
                  Button(
                    icon: SvgPicture.asset(
                      "assets/status.svg",
                      width: 36,
                      height: 36,
                    ),
                    label: "Status",
                    onPressed: () {
                      context
                          .read<ApiBloc>()
                          .add(SnapRaidSyncEvent("snapraid/status"));
                    },
                  ),
                  Button(
                    icon: SvgPicture.asset(
                      "assets/diff.svg",
                      width: 36,
                      height: 36,
                    ),
                    label: "Diff",
                    onPressed: () {
                      context
                          .read<ApiBloc>()
                          .add(SnapRaidSyncEvent("snapraid/diff"));
                    },
                  ),
                  Button(
                    icon: SvgPicture.asset(
                      "assets/sync.svg",
                      width: 36,
                      height: 36,
                    ),
                    onPressed: () {
                      context
                          .read<ApiBloc>()
                          .add(SnapRaidSyncEvent("snapraid/sync"));
                    },
                    label: "Sync",
                  ),
                  Button(
                    onPressed: () {
                      context
                          .read<ApiBloc>()
                          .add(SnapRaidSyncEvent("snapraid/scrub"));
                    },
                    icon: SvgPicture.asset(
                      "assets/scrub.svg",
                      width: 36,
                      height: 36,
                    ),
                    label: "Scrub",
                  ),
                ],
              ),
            ),
            Padding(
              padding: const EdgeInsets.all(30.0),
              child: Row(
                mainAxisAlignment: MainAxisAlignment.spaceBetween,
                children: [
                  Button(
                    onPressed: () {
                      context
                          .read<ApiBloc>()
                          .add(SnapRaidSyncEvent("snapraid/list"));
                    },
                    label: "List",
                    icon: SvgPicture.asset(
                      "assets/list.svg",
                      width: 36,
                      height: 36,
                    ),
                  ),
                  Button(
                    onPressed: () {
                      context
                          .read<ApiBloc>()
                          .add(SnapRaidSyncEvent("snapraid/dup"));
                    },
                    label: "Dup",
                    icon: SvgPicture.asset(
                      "assets/dup.svg",
                      width: 36,
                      height: 36,
                    ),
                  ),
                  Button(
                    onPressed: () {
                      context
                          .read<ApiBloc>()
                          .add(SnapRaidSyncEvent("snapraid/smart"));
                    },
                    icon: SvgPicture.asset(
                      "assets/smart.svg",
                      width: 36,
                      height: 36,
                    ),
                    label: "Smart",
                  ),
                  Button(
                    onPressed: () {
                      context
                          .read<ApiBloc>()
                          .add(SnapRaidSyncEvent("snapraid/check"));
                    },
                    icon: SvgPicture.asset(
                      "assets/check.svg",
                      width: 36,
                      height: 36,
                    ),
                    label: "Check",
                  ),
                ],
              ),
            ),
            Expanded(
              child: BlocBuilder<WebsocketBloc, WebsocketApiStartState>(
                  builder: ((context, state) {
                return StreamBuilder(
                    stream: state.channel.stream,
                    builder: ((context, snapshot) {
                      if (snapshot.hasData) {
                        final json = jsonDecode(snapshot.data);
                        final msgType = json["MsgType"];
                        if (msgType == 1) {
                          loglist.Append(json["Message"]);
                        } else if (msgType == 0) {
                          loglist.Clear();
                        }
                      }
                      return loglist;
                    }));
              })),
            ),
          ],
        ),
      ),
    );
  }
}
