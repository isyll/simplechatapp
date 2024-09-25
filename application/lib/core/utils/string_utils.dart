import '../extensions/string_extension.dart';

/// Converts snake_case string to camelCase.
String snakeToCamelCase(String snakeCase) {
  if (snakeCase.trim().isEmpty) return snakeCase;
  final words = snakeCase.split('_');

  for (int i = 0; i < words.length; i++) {
    words[i] = words[i].isNotEmpty
        ? words[i][0].toUpperCase() + words[i].substring(1)
        : '';
  }

  return words.join('').lcfirst();
}

/// Converts camelCase string to snake_case.
String camelToSnakeCase(String camelCase) {
  if (camelCase.trim().isEmpty) return camelCase;

  final regex = RegExp(r'(?<=[a-z])([A-Z])|(?<=[A-Z])([A-Z][a-z])');
  final snakeCase = camelCase.replaceAllMapped(
      regex, (Match match) => '_${match.group(0)!.toLowerCase()}');

  return snakeCase;
}
