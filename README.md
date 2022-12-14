# Go basic
## build, install, run 차이점
- build는 go 소스코드를 가지고 실행 가능한 파일(exe)로 생성해줌. (테스트, 컴파일용 명령어)
- install은 pkg(외부 라이브러리) 등을 모두 합쳐서 하나의 application이라 생각하고, binary 파일로 만들어 줌.
- run은 간단하게 터미널에서 실행해볼 수 있는 명령어

## Go 장점 및 특징
- 간결한 문법 및 단순함
- 병행 프로그래밍 지원
  - 타 언어에서는 쓰레드를 사용하여 멀티 쓰레드, 즉 병렬 프로그래밍을 가능케함.
  - 하지만 고에선 고루틴이 지원되기 때문에 따로 쓰레드를 생성할 필요 없이 병행 프로그래밍 가능.
- 정적 타입 및 동적 실행
- 간편한 협업 지원
- 컴파일 및 실행속도 빠름
- 제네릭 및 예외 처리 미지원
- 컨벤션 통일

### 변수 var, 상수 const
- 변수에서 shortVariable 기억! => 스택 {} 안 스코프를 가진다.
- 여러 개를 선언하고 다른 타입을 할당할 수 있음.
- 다른 언어들과 마찬가지로 const 는 선언과 동시에 초기화 필요, 재할당 X

### 제어문

#### 1. if
- if, if else, if else if 가능. 중괄호 위치 지정, 생략 불가능

#### 2. switch
- switch 뒤 Expression 생략 가능
- case 뒤 Expression 사용 가능
- 자동 break 때문에 fallthrough 존재
- type 분기 -> 값이 하닌 변수 Type으로 분기 가능

#### 3. for
- go에서 제공하는 유일한 반복문이다.
- 다른 프로그래밍 언어에서 while 과 비슷하게 사용 가능.

### Go 문법 특징 정리
- 주석 : // /**/
- 세미콜론
  - 컴파일러가 자동으로 세미콜론 추가하므로 추가하면 에러
- 연산[증감, 전치]
  - 전치 연산자 (++i) 비허용, 후치 연산자 (i++) 반환 값 없음.
- 코드 자동 서식
  - gofmt가 자동으로 파일의 포맷팅을 해줌.


### package
- Go의 지향점은 하나의 모듈은 응집도가 높고 결합도가 낮은, 즉 우리가 지향점으로 삼고 있는 소프트웨어 아키텍쳐와 비슷하다고 할 수 있다.
- 하나의 package 는 하나의 기능을 표출할 수 있어야함.
- 작은 패키지를 결합해서 프로그램을 작성 할 것을 권고하고 있다.
**네이밍 규칙**
- 대문자 public, 소문자 private 접근 제어를 가짐
- 별칭(alias) 패키지 사용 가능.  
**초기화 메서드**
- `func init() {}` : main보다 먼저 호출됨.
- package import 하면, init method가 있다면 호출된다.
- [https://go.dev/doc/effective_go#initialization](https://go.dev/doc/effective_go#initialization)

### 데이터 타입
#### Boolean 
- 암묵적 형 변환 X : 0, Nil -> false 변환 없음.
- 논리 연산자 &&, ||, ! 연산 가능
- 논리 연산자 >, <, !=, ==, <=, >= 

### Numberic
- 정수, 실수, 복소수
- 32bit, 64bit, unsigned(양수)
- 정수: 8진수(0), 16진수(0x), 10진수 => 시작하는 prefix가 0이면 8진수, 0x이면 16진수이다.
- rune(유니코드), byte(아스키) 정수 할당 가능
- 실수(부동소수점)
  - float32(7자리), float64(15자리)
    - 32 -> 64 변환할 때 주의!
- 같은 크기의 데이터 타입을 연산할 때는 캐스팅 후 해야할 것.
  - 같은 타입만 연산 가능
- 비트 오버플로우 유의!


### String
- 큰 따옴표 "", 백스쿼트 ``
- 문자 char 타입 존재하지 않음. -> rune(int32)로 문자 코드 값으로 표현
- 문자 : ''로 작성
- 자주 사용하는 escape : \\\ , \\`, \\"
- len 함수는 바이트를 리턴해준다.
- len 배열을 넣으면 길이를 리턴해주기 때문에, 한글 스트링을 rune 배열로 만들고 호출하면 길이를 리턴해준다.
#### 문자열 표현
- UTF-8 인코딩 (유니코드 문자 집합) -> 바이트 수 주의!
#### 문자열 연산
- 일반 + 연산도 가능하지만, 배열을 통해 append하고 Join을 이용하는 것이 유리


## data structure
### array
- 값 타입이다
- 예를 들어 `arr2 := arr1` 로 할당시키면 다른 주소에 배열을 복사한다.
- 배열은 용량, 길이가 항상 같음


### 배열 슬라이스 비교
- 길이고정 vs 길이가변
- 값 타입 vs 참조 타입
- 복사 전달 vs 참조 값 전달
- 전체 비교 연산자 사용 가능 vs 비교 연산자 사용 불가
- 대부분 슬라이스 사용한다.

### slice
- 참조 타입이다.
- []타입, []타입{초기화} 로 선언 및 할당 가능
- make 함수로 할당 가능
- 부분적으로 슬라이스 추출은 참조 타입으로써, 원본 데이터가 변경되니 주의.

### map
- 래퍼런스 타입 (참조 값 전달)
- 비교 연산자 사용 불가능(참조 타입이므로)
- 특징 : 참조타입(Key)로 사용 불가능, 값(Value)으로 모든 타입 사용 가능
- make 함수 및 축약(리터럴)로 초기화 가능
- 순서가 없다.
- 없는 키를 조회하면 default 초기화 값을 리턴해주기 때문에 두 번째 리턴 값으로 키 존재 유무 확인해야함.

### pointer
- 포인터지원 X(파이썬, C#, JAVA 등)
- 주소 값은 직접 변경 불가능
- *(애스터리스크)로 사용
- nil로 초기화 (nil == 0)
```go
var a *int // 포인터형 변수 선언, 당연히 a의 메모리 주소도 가지고 있으며, 다른 변수의 메모리 주소를 할당.
var a *int = new(int) 

```

### function
- 선언 : func 키워드로 선언
- func 함수명(매개변수) (반환타입 or 반환 값 변수명) : 반환 값 존재
- func 함수명() (반환타입 or 반환 값 변수명) : 매개 변수 없음, 반환 값 존재
- func 함수명(매개변수) : 매개변수 존재, 반환 값 없음
- func 함수명() : 매개변수 없음, 반환 값 없음
- 타 언어와 달리 반환 값(return value) 다수 가능

### defer
- Defer 함수 실행 (지연)
- Defer를 호출한 함수가 종료되기 직전에 호출된다.
- 타 언어의 Finally 문과 비슷
- 주로 리소스 반환, 열린 파일 닫기, Mutex 잠금 해제

### closure
- 익명함수 사용할 경우 함수를 변수에 할당해서 사용가능
- 함수 안에서 함수를 선언 및 정의 가능 -> 이 때 외부 함수에 선언 된 변수에 접근 가능한 함수
- 함수가 선언되는 순간에 함수가 실행 될 때 실체의 외부 변수에 접근하기 위한 스냅샷(객체)
- 함수를 호출할 때 이전에 존재했던 값을 유지하기 위해서 -> 비동기, 누적카운트, 무분별한 전역변수 남발...
- 남발 -> 객체들이 메모리에 존재하므로, -> 메모리 부족, 오버플로우 현상, 자원을 무분별하게 사용 가능성
- 클로저 정확하게 이해하고 사용
- 예제 코드
```go
m, n := 5, 10            // 변수가 캡쳐 (선언과 동시에 캡쳐됨.)
sum := func(c int) int { //익명함수 선언과 동시에 할당
  return m + n + c // 지역변수 소멸 되지 않음. (함수 호출시마다 사용 가능)
}
// 캡쳐됐다는 말은 변수의 메모리 주소를 함수를 선언할 때 할당한다.(캡쳐한다)
```

## 사용자 정의 타입
- Go는 객체 지향 타입을 구조체로 정의한다 (클래스, 상속 개념 없음)
- 객체지향 -> 클래스(속성 : 멤버변수, 기능(상태 : 메소드)) : 코드의 재사용성, 코드의 관리 용이, 신뢰성이 높은 프로그래밍
- Go는 전형적인 객체지향의 특징을 가지고 있지 않지만, 인터페이스 -> 다형성 지원, 구조체를 클래스 형태의 코딩 가능
- 객체지향의 기본 개념 -> Go에서 포함하고 있다. -> 객체 지향 프로그래밍 언어이다
- 상태, 메소드 분리해서 정의(결합성 없음)
- 사용자 정의 타입 : 구조체, 인터페이스, 기본타입(int, float, string), 함수
- 구조체와 -> 메소드 연결을 통해서 타 언어의 클래스 형식처럼 사용 가능하다!(객체지향)
- 리시버(구조체 메소드) 전달(값, 참조) 형식
- 함수는 기본적으로 값 호출 -> 변수의 값이 복사 후 내부 전달(원본 수정X) -> 맵, 슬라이스 등은 참조 전달
- 리시버(구조체)도 마찬가지로 포인터를 활용해서 메소드내에서 원본 수정 가능
  - 반대로 말하면, 값 전달으로는 원본 값 수정이 되지 않음 (객체의 상태 수정 X)

### 생성자 관례
```go
type Account struct {
	number   string
	balance  float64
	interest float64
}

func NewAccount(number string, balance float64, interest float64) *Account {
	return &Account{number, balance, interest}
}
```

### 구조체 임베디드 패턴
- 다른 관점으로 메소드를 재사용하는 장점 제공
- 상속을 허용하지 않는 Go 언어에서 메소드 상속 활용을 위한 패턴
```go
type Employee struct {
	name   string
	salary float64
	bonus  float64
}

func (e Employee) Calculate() float64 {
	return e.salary + e.bonus
}

type Executives struct {
	Employee
	specialBonus float64
}
```

### 구조체 임베디드 메소드 오버라이딩 패턴
- 직접 구조체 타입을 명시해서 리시버를 임베디드 구조체와 이름을 같게 한다면 Go에서 오버라이딩 기능 제공 

## 인터페이스
- 객체의 동작을 표현, 골격
- 단순히 동작에 대한 방법만 표시
- 추상화 제공
- 인터페이스 메소드를 구현한 타입은 인터페이스로 사용 가능
- Go 언어를 유연하게 사용 가능
- 덕타이핑 : Go 언어의 독창적인 특성
- 인터페이스 -> 자바(전략 패턴, 템플릿 메소드, 팩토리 메소드 패턴, 어댑터 패턴 .....)
- 디자인 패턴 측면에서 client 입장은 메소드의 구체적인 클래스를 몰라도 인터페이스 정의된 메소드를 사용하는 객체를 보장한다.
- 클래스간의 결합도 감소 -> 유지보수성 향상, 기능 추가의 용이성 -> 독립적인 프로그래밍 가능

## 고루틴
- 고루틴(Goroutine)
- 코어 수에 상관없이 동시에 실행되는 것처럼 보이는 것
- 실제로는 코어 수에 비례해서 동시에 실행됨
- 비동기적으로 동작하는 함수 루틴
- 채널을 통한 통신으로 동기화 가능
- 멀티 코어 CPU에서 각 코어는 별도의 고루틴을 동시에 실행
- 고루틴은 메인 함수가 종료되어도 계속 실행됨
- 주의 : 고루틴은 메인 함수가 종료되면 같이 종료됨
- 주의 : 고루틴은 동시에 실행되는 것처럼 보이지만 실제로는 순차적으로 실행됨
- 주의 : 고루틴은 메모리 공유 시에 문제 발생 -> 고루틴 동기화 필요
- 주의 : 고루틴은 많은 메모리를 사용 -> 고루틴 사용 시 메모리 사용량 고려

- 멀티 쓰레드 장점과 단점
- 장점 : 응답성이 좋아짐, 자원 공유가 용이, 작업이 분리되어 코드 간결
- 단점 : 구현하기 어려움(복잡성이 증가), 테스트 및 디버깅 어려움, 동기화 처리 필요, 자원 공유 시 문제 발생 가능성 증가
- 고루틴 장점과 단점
- 장점 : 매우 가볍고 간단, 자원 공유가 용이
- 단점 : 응답성이 떨어짐, 자원 공유 시 문제 발생 가능성 증가

## 채널(Channel)
- 고루틴 간에 통신을 위한 메커니즘
- 고루틴간의 상호 정보(데이터) 교환 및 실행 흐름 동기화 위해서 사용 : 채널(동기식, 래퍼런스 타입)
- 실행 흐름 제어 가능(동기, 비동기) -> 일반 변수로 선언 후 사용 가능
- 데이터 전달 자료형 선언 후 사용(지정 된 타입만 주고 받을 수 있음)
- interface{} 전달을 통해서 자료형 상관없이 전송 및 수신가능
- 값을 전달(복사 후 : bool, int등), 포인터(슬라이스, 맵) 등을 전달시에는 주의 > 동기화 사용(Mutex)
- 멀티 프로세싱 처리에서 교착상태(경쟁) 주의
- <-, -> (채널 <- 데이터 : 송신), ( <- 채널 : 수신)

### 고루틴 동기화 고급
- Once : 한 번만 실행(주로 초기화에 사용)
- Do로 실행
- 대기 그룹
- 실행 흐름 변경(고루틴이 종료 될 때까지 대기 가능)
- WaitGroup : Add(고루틴 추가), Done(작업 종료 알림), Wait(고루틴 종료시까지 대기)
- Add로 추가 된 고루틴 개수와 Done으로 종료되는 알림 횟수가 같아야 정확하게 동작한다.
- 원자성 사용 -> 기능적으로 분할 불가능한 완전 보증된 일련의 조작, 모두 성공하거나, 모두 실패
- 모든 조작이 완료 될 때까지 다른 프로세스 개입 불가
- sync/atomic 패키지에서 원자적 연산자 제공
- [https://golang.org/pkg/sync/atomic/](https://golang.org/pkg/sync/atomic/)
- 주로 공용 변수에 관한 계산 사용

## 에러 처리
- 소프트웨어의 품질 향상시 가장 중요한 것 -> 유형 코드 및 에러 정보 등 정보를 남기는 것!
- Go에서는 기본적으로 error 패키지에서 error 인터페이스를 제공한다.
- type error interface { Error() string }
- 즉, Error 메서드를 구현하면 사용자 정의 에러 처리 제작 가능.
- 기본적으로 메서드마다 리턴 타입 2개(리턴 값, 에러)
- 주로 
  1. Errorf(**에러 타입 리턴**) 메소드
  2. Fatal(**프로그램 종료**) 메소드를 통해서 에러 출력
### 에러 처리 고급
- error 타입이 아닌 경우 에러 처리 방법?
- Error 메소드를 구현해서 사용자 정의 에러 처리.
- 구조체를 사용해서 세부적인 정보 출력 가능
- 덕타입을 활용.
### Panic, Recover
- Panic : 프로그램을 비정상적으로 종료
- Recover : 비정상 종료를 복구
### Go Panic 함수
- 사용자가 에러 발생 가능
- Panic 함수는 호출 즉시, 해당 메소드를 즉시 중지시키고 defer 함수를 호출하고 자기 자신을 호출한 곳으로 리턴
- 런타임 이외에 사용자가 코드 흐름에 따라 에러를 발생 시킬 때 중요!
- 문법적인 에러는 아니지만, 논리적인 코드 흐름에 따른 에러 발생 처리 가능

### Go Recover 함수
- 에러 복구 가능
- Panic으로 발생한 에러를 복구 후 코드 재실행(프로그램 종료되지 않는다)
- 즉, 코드 흐름을 정상 상태로 복구 하는 기능
- Panic에서 설정한 메시지를 받아올 수 있다.

## 파일 입출력
### 파일 쓰기
- Create : 새 파일 작성 및 파일 열기
- Close : 리소스 닫기
- Write, WriteString, WriteAt : 파일 쓰기
- 각 운영체제 권한 주의

### 파일 읽기
- Open : 기존 파일 열기
- Close : 리소스 닫기
- Read, ReadAt : 파일 읽기
- 각 운영체제 권한 주의(오류메시지 확인)
- 탐색 Seek 중요
- [https://golang.org/pkg/os](https://golang.org/pkg/os)

### 파일 읽기, 쓰기 -> ioutil 패키지 활용
- 더욱 편리하고 직관적으로 파일을 읽고 쓰기 가능
- 아래 메소드 확인 가능
- WriteFile(), ReadFile(), ReadAll() 등 사용 가능
- [https://golang.org/pkg/io/ioutil](https://golang.org/pkg/io/ioutil) => Deprecated
### 파일 읽기, 버퍼사용 쓰기 -> bufio 패키지 사용
- ioutil, bufio 등은 io.Reader, io.Writer 인터페이스를 구현함 -> 즉 Writer, Read 메소드를 사용 가능
```go
type Reader interface {
    Read(p []byte) (n int, err error)
}
type Writer interface {
    Write(p []byte) (n int, err error)
}
```
- 즉, bufio의 NewReader, NewWriter를 통해서 객체 반환 후 메소드 사용
```text
상태 
a ----> a
b ----> ab
c ----> abc
d ----> abcd
e ----> e       ----> abcd
e ----> ef      ----> abcd
e ----> efg     ----> abcd
e ----> efgh    ----> abcd
e ----> i       ----> abcdefgh
```
- 버퍼를 사용하면, 파일을 읽고 쓰는 횟수를 줄일 수 있다.

## 사용자 패키지 작성 & Go 문서화
- 기준은 GOPATH/src
- 패키지 폴더명(디렉토리명) 명확하게 지정
- 패키지 파일의 package 이름으로 사용한다. 또는 alias 사용
- package main 을 제외하고 package 문서에 등록
- 기본적으로 GOROOT의 패키지(pkg) 검색 -> GOPATH의 패키지(src/pkg) 검색
- go install 명령어 실행의 경우 -> GOPATH/pkg에 등록 (문서화)
- godoc -http=:6060(임의의포트) -> pkg 이동 -> 본인 패키지 메소드 및 주석 확인(패키지, 타입, 메소드) 주석
### 외부 저장소 패키지 설치
2가지 방법
1. import 선언 후 폴더 이동 후 go get 설치 -> 사용
2. go get 패키지 주소 설치 -> 선언