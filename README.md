# golang_network

## command

`go run main.go`

- compile and run

`go build main.go`

- compile go

`go clean main.go`

- cleaning object and cache

`go mod init ${module_name}`

- 기본적으로 go는 GOROOT 에 있는 lib 만 쓰는데 나만의 커스텀을 쓰기 위해서 `go.mod`가 필요한데 이걸 만들때 사용함.

## string formatt

- %v : 구조체의 값을 출력한다.
- %+ek. : 구조체의 필드이름과 값을 출력한다.
- %#g : 호출된 함수이름과 구조체의 이름 같은 소스코드 정보까지 함께 출력한다.
- %T : 타입을 출력한다.
- %t : 불리언의 값을 true 혹은 false 문자열로 출력한다.
- %d : Integer 값을 출력하기 위해서 사용한다. 10자리 크기의 정수를 포함한 문자열 형식을 가진다.
- %b : 이진(바이너리)값을 출력한다.
- %c : 정수에 해당하는 문자를 출력한다.
- %x : Hex 인코딩 값을 출력한다.
- %f : 부동 소숫 점 값을 출력한다.
- %e : 과학적 표기법으로 출력한다.
- %E : 과학적 표기법으로 출력한다. %e와 다른점은 e가 대문자인지 소문자인지
- %s : 문자열을 출력할 때 사용한다. %d와 더불어 가장 많이 사용하는 옵션
- %q : 문자열에 있는 쌍다움표를 그대로 출력한다.
- %x : 먼저 값을 integer로 변환한 다음 16비트 문자열로 변환해서 출력한다.
- %p : 포인터의 주소 값을 출력한다.
- %nd : 포맷팅에 사용하는 숫자 n을 이용해서 출력 할 넓이를 설정할 수 있다. %6d인 경우 6칸의 넓이를 가진다. 채우지 못한 곳은 왼쪽 부터 스페이스 문자로 채워진다.
- %-nd : %nd와 달리 왼쪽부터 데이터가 채워지고 나머지 공간을 스페이스 문자가 채운다.

## Pacakge Export

- 대문자 시작은 외부 접근 가능
- 소문자 시작은 외부 접근 불가
