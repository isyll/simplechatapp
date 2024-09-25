class Failure extends Error {}

class UserNotFound extends Failure {}

class InvalidFieldValue extends Failure {
  final String fieldName;

  InvalidFieldValue({
    required this.fieldName,
  });
}

class UserNotLogged extends Failure {}
