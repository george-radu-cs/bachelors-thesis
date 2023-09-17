class CardReview {
  // ...
  final Duration timeToAnswer;
  final DateTime reviewTimestamp;

  factory CardReview.fromMap(Map<String, dynamic> json) => CardReview(
    // ...
    timeToAnswer: Duration(milliseconds: json['timeToAnswer']),
    reviewTimestamp: DateTime.fromMillisecondsSinceEpoch(json['reviewTimestamp']),
  );
  
  Map<String, dynamic> toMap() => {
    // ...
    'timeToAnswer': timeToAnswer.inMilliseconds,
    'reviewTimestamp': reviewTimestamp.millisecondsSinceEpoch,
  };
}

class Settings {
  // ...
  final bool showOCRDialog;

  factory Settings.fromMap(Map<String, dynamic> json) => Settings(
    // ...
    showOCRDialog: json['showOCRDialog'] == 1,
  );

  Map<String, dynamic> toMap() => {
    // ...
    'showOCRDialog': showOCRDialog ? 1 : 0,
  };
}