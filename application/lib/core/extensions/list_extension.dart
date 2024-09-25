import 'dart:math';

extension ListExtension<E> on List<E> {
  E get randomItem => this[Random().nextInt(length)];

  E? firstWhereOrNull(bool Function(E element) test) {
    for (E element in this) {
      if (test(element)) {
        return element;
      }
    }
    return null;
  }

  List<T> mapWithIndex<T>(T Function(int index, E item) callback) {
    return asMap()
        .entries
        .map((entry) => callback(entry.key, entry.value))
        .toList();
  }
}
