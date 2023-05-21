# go-error-wrap

## 概要

- Go言語でerrorをwrapした際に、wrap 元と比較する実装を学ぶ
    - 自作エラーに`Unwrap()`を明示的に実装する
    - `errors.Is`と`errors.As`は wrap されていても途中で操作を切り上げるため