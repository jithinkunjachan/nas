import 'package:flutter/material.dart';

class TextArea extends StatefulWidget {
  final textArea = _TextArea();

  TextArea({super.key});

  void Write(String string) {
    textArea.Write(string);
  }

  void Clear() {
    textArea.Clear();
  }

  @override
  State<StatefulWidget> createState() {
    return textArea;
  }
}

class _TextArea extends State<TextArea> {
  final buffer = StringBuffer();

  void Write(String value) {
    setState(() {
      buffer.write("\n");
      buffer.write(value);
    });
  }

  void Clear() {
    setState(() {
      buffer.clear();
    });
  }

  @override
  Widget build(BuildContext context) {
    return Container(
      height: 400,
      child: Scrollbar(
        child: SingleChildScrollView(
          scrollDirection: Axis.vertical,
          child: Text(
            buffer.toString(),
            style: const TextStyle(fontSize: 6),
          ),
        ),
      ),
    );
  }
}
