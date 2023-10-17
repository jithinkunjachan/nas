import 'package:flutter/material.dart';

class LogList extends StatefulWidget {
  final loglist = _LogList();

  LogList({super.key});

  void Append(String string) {
    loglist.Append(string);
  }

  void Clear() {
    loglist.Clear();
  }

  @override
  State<StatefulWidget> createState() {
    return loglist;
  }
}

class _LogList extends State<LogList> {
  final List<String> list = [];

  void Append(String value) {
    setState(() {
      list.add(value);
    });
  }

  void Clear() {
    setState(() {
      list.clear();
    });
  }

  @override
  Widget build(BuildContext context) {
    return ListView.builder(
      itemCount: list.length,
      itemBuilder: (context, index) {
        return Text(list[index]);
      },
    );
  }
}
