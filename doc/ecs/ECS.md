# ECS (Entity Component System)

## 개요

ECS는 **Entity Component System**의 약자로, 게임 개발에서 자주 사용되는 아키텍처 패턴입니다.
전통적인 객체 지향의 상속 구조 대신 **조합(Composition)**을 사용하여 유연하고 확장 가능한 게임 구조를 만듭니다.

## 핵심 구성 요소

### 1. Entity (엔티티)

- **정의**: 게임 세계에 존재하는 객체를 나타내는 고유 식별자 (ID)
- **특징**:
  - 데이터나 로직을 직접 가지지 않음
  - 단순히 컴포넌트들을 묶는 컨테이너 역할
  - 일반적으로 정수형 ID로 표현

**예시**:
```go
type Entity uint64

var playerEntity Entity = 1
var enemyEntity Entity = 2
var bulletEntity Entity = 3
```

### 2. Component (컴포넌트)

- **정의**: 순수한 데이터만을 담고 있는 구조체
- **특징**:
  - 로직(메서드)을 포함하지 않음
  - 재사용 가능한 데이터 덩어리
  - 엔티티의 특성을 정의

**일반적인 컴포넌트 예시**:

```go
// 위치 컴포넌트
type Position struct {
    X float64
    Y float64
}

// 속도 컴포넌트
type Velocity struct {
    DX float64
    DY float64
}

// 스프라이트 컴포넌트
type Sprite struct {
    Texture string
    Width   int
    Height  int
}

// 체력 컴포넌트
type Health struct {
    Current int
    Max     int
}

// 충돌 컴포넌트
type Collision struct {
    Width  float64
    Height float64
}
```

### 3. System (시스템)

- **정의**: 특정 컴포넌트 조합을 가진 엔티티들을 처리하는 로직
- **특징**:
  - 데이터(컴포넌트)와 분리된 순수 로직
  - 매 프레임마다 실행되거나 이벤트 기반으로 실행
  - 특정 컴포넌트 조합을 필터링하여 처리

**시스템 예시**:

```go
// 이동 시스템: Position + Velocity를 가진 모든 엔티티 처리
type MovementSystem struct{}

func (s *MovementSystem) Update(dt float64) {
    // Position과 Velocity를 모두 가진 엔티티 탐색
    for _, entity := range GetEntitiesWithComponents(Position{}, Velocity{}) {
        pos := GetComponent[Position](entity)
        vel := GetComponent[Velocity](entity)

        pos.X += vel.DX * dt
        pos.Y += vel.DY * dt
    }
}

// 렌더링 시스템: Position + Sprite를 가진 모든 엔티티 처리
type RenderSystem struct{}

func (s *RenderSystem) Update() {
    for _, entity := range GetEntitiesWithComponents(Position{}, Sprite{}) {
        pos := GetComponent[Position](entity)
        sprite := GetComponent[Sprite](entity)

        DrawSprite(sprite.Texture, pos.X, pos.Y)
    }
}

// 충돌 시스템: Position + Collision을 가진 모든 엔티티 처리
type CollisionSystem struct{}

func (s *CollisionSystem) Update() {
    entities := GetEntitiesWithComponents(Position{}, Collision{})

    for i := 0; i < len(entities); i++ {
        for j := i + 1; j < len(entities); j++ {
            if CheckCollision(entities[i], entities[j]) {
                HandleCollision(entities[i], entities[j])
            }
        }
    }
}
```

## ECS의 장점

### 1. 유연성
- 상속 계층 없이 컴포넌트 조합으로 다양한 엔티티 생성
- 런타임에 컴포넌트 추가/제거 가능

```go
// 예: 일반 적 -> 보스 적으로 변환
AddComponent(enemy, Health{Current: 1000, Max: 1000})
AddComponent(enemy, BossAI{})
```

### 2. 재사용성
- 컴포넌트와 시스템을 여러 엔티티에서 재사용
- 중복 코드 최소화

### 3. 성능
- 데이터 지향 설계로 캐시 효율성 향상
- 같은 타입의 컴포넌트를 연속된 메모리에 배치 가능
- 병렬 처리에 유리

### 4. 확장성
- 새로운 기능을 컴포넌트와 시스템으로 쉽게 추가
- 기존 코드 수정 없이 확장 가능

### 5. 디버깅 용이성
- 데이터와 로직이 분리되어 추적이 쉬움
- 특정 시스템만 켜고 끌 수 있음

## Tetris 게임에 적용하기

### 엔티티 예시

```go
// 테트리스 블록 (테트로미노)
tetromino := CreateEntity()

// 게임 보드
board := CreateEntity()

// 다음 블록 프리뷰
nextPiece := CreateEntity()

// UI 요소
scoreText := CreateEntity()
```

### 컴포넌트 예시

```go
// 테트리스 블록 모양 컴포넌트
type Shape struct {
    Type     string  // I, O, T, L, J, S, Z
    Rotation int     // 0, 90, 180, 270
    Blocks   [][]int // 블록 배열
}

// 그리드 위치 컴포넌트
type GridPosition struct {
    X int
    Y int
}

// 색상 컴포넌트
type Color struct {
    R, G, B, A uint8
}

// 낙하 속도 컴포넌트
type FallSpeed struct {
    Speed        float64 // 블록/초
    Timer        float64
    FastFalling  bool
}

// 게임 보드 컴포넌트
type Board struct {
    Width  int
    Height int
    Grid   [][]int // 고정된 블록 정보
}

// 점수 컴포넌트
type Score struct {
    Value int
    Level int
    Lines int
}
```

### 시스템 예시

```go
// 낙하 시스템
type FallSystem struct{}

func (s *FallSystem) Update(dt float64) {
    for _, entity := range GetEntitiesWithComponents(GridPosition{}, FallSpeed{}) {
        pos := GetComponent[GridPosition](entity)
        fall := GetComponent[FallSpeed](entity)

        fall.Timer += dt
        speed := fall.Speed
        if fall.FastFalling {
            speed *= 10
        }

        if fall.Timer >= 1.0/speed {
            pos.Y++
            fall.Timer = 0
        }
    }
}

// 충돌 감지 시스템
type CollisionSystem struct{}

func (s *CollisionSystem) Update() {
    for _, entity := range GetEntitiesWithComponents(GridPosition{}, Shape{}) {
        if CheckBoardCollision(entity) {
            LockPiece(entity)
            SpawnNewPiece()
        }
    }
}

// 회전 시스템
type RotationSystem struct{}

func (s *RotationSystem) HandleInput(input Input) {
    if input.RotatePressed {
        for _, entity := range GetEntitiesWithComponents(Shape{}, GridPosition{}) {
            shape := GetComponent[Shape](entity)
            shape.Rotation = (shape.Rotation + 90) % 360
            UpdateShapeBlocks(shape)
        }
    }
}

// 라인 클리어 시스템
type LineClearSystem struct{}

func (s *LineClearSystem) Update() {
    for _, entity := range GetEntitiesWithComponents(Board{}, Score{}) {
        board := GetComponent[Board](entity)
        score := GetComponent[Score](entity)

        clearedLines := CheckAndClearLines(board)
        if clearedLines > 0 {
            score.Lines += clearedLines
            score.Value += CalculateScore(clearedLines)
        }
    }
}

// 렌더링 시스템
type RenderSystem struct{}

func (s *RenderSystem) Update() {
    // 보드 렌더링
    for _, entity := range GetEntitiesWithComponents(Board{}) {
        DrawBoard(entity)
    }

    // 테트로미노 렌더링
    for _, entity := range GetEntitiesWithComponents(GridPosition{}, Shape{}, Color{}) {
        DrawTetromino(entity)
    }
}
```

### 게임 루프 예시

```go
func GameLoop() {
    systems := []System{
        &InputSystem{},
        &RotationSystem{},
        &FallSystem{},
        &CollisionSystem{},
        &LineClearSystem{},
        &RenderSystem{},
    }

    for !gameOver {
        dt := CalculateDeltaTime()

        for _, system := range systems {
            system.Update(dt)
        }
    }
}
```

## 전통적인 OOP vs ECS

### 전통적인 OOP
```go
type GameObject struct {
    position Vector2
    velocity Vector2
}

type Player struct {
    GameObject  // 상속
    health int
    sprite Sprite
}

type Enemy struct {
    GameObject  // 상속
    health int
    sprite Sprite
    ai     AI
}

// 문제점: 코드 중복, 깊은 상속 계층, 유연성 부족
```

### ECS 접근법
```go
// 엔티티는 단순히 ID
player := CreateEntity()
AddComponent(player, Position{})
AddComponent(player, Velocity{})
AddComponent(player, Health{})
AddComponent(player, Sprite{})
AddComponent(player, PlayerInput{})

enemy := CreateEntity()
AddComponent(enemy, Position{})
AddComponent(enemy, Velocity{})
AddComponent(enemy, Health{})
AddComponent(enemy, Sprite{})
AddComponent(enemy, AI{})

// 장점: 조합 가능, 재사용 가능, 유연함
```

## ECS 구현 방식

ECS를 구현하는 방법은 여러 가지가 있습니다. 각 방식은 성능, 메모리 효율성, 구현 복잡도에서 트레이드오프가 있습니다.

### 핵심 개념

엔티티가 컴포넌트를 **직접 소유하지 않습니다**. 대신:
- ✅ World/Registry가 컴포넌트 타입별 저장소를 관리
- ✅ 엔티티는 단순히 ID (정수)
- ✅ 맵이나 배열로 `Entity ID → Component` 매핑

### 1. Component Table 방식 (가장 일반적)

컴포넌트 타입별로 별도의 맵을 만들고, 엔티티 ID를 키로 사용하는 방식입니다.

```go
// World가 컴포넌트들을 타입별로 관리
type World struct {
    // 컴포넌트 타입별로 맵 보관
    positions  map[Entity]Position
    velocities map[Entity]Velocity
    sprites    map[Entity]Sprite
    healths    map[Entity]Health

    nextEntityID Entity
}

// 엔티티 생성
func (w *World) CreateEntity() Entity {
    entity := w.nextEntityID
    w.nextEntityID++
    return entity
}

// 컴포넌트 추가
func (w *World) AddPosition(entity Entity, pos Position) {
    w.positions[entity] = pos
}

func (w *World) AddVelocity(entity Entity, vel Velocity) {
    w.velocities[entity] = vel
}

// 컴포넌트 가져오기
func (w *World) GetPosition(entity Entity) (Position, bool) {
    pos, exists := w.positions[entity]
    return pos, exists
}

// 특정 컴포넌트를 가진 엔티티들 찾기
func (w *World) Query() []Entity {
    var entities []Entity
    for entity := range w.positions {
        if _, hasVel := w.velocities[entity]; hasVel {
            entities = append(entities, entity)
        }
    }
    return entities
}

// 시스템 예시
type MovementSystem struct{}

func (s *MovementSystem) Update(w *World, dt float64) {
    for _, entity := range w.Query() {
        pos, _ := w.GetPosition(entity)
        vel, _ := w.GetVelocity(entity)

        pos.X += vel.DX * dt
        pos.Y += vel.DY * dt

        w.AddPosition(entity, pos)
    }
}
```

**장점**:
- 구현이 간단하고 직관적
- 타입 안전성 보장
- 디버깅이 쉬움

**단점**:
- 컴포넌트 타입마다 맵을 만들어야 함
- 쿼리 성능이 O(n)
- 제네릭 코드 작성이 어려움

### 2. Go 제네릭 사용 (타입 안전 + 유연성)

Go 1.18+의 제네릭을 활용하면 더 유연하게 구현할 수 있습니다.

```go
type World struct {
    components map[reflect.Type]map[Entity]interface{}
    nextEntityID Entity
}

func NewWorld() *World {
    return &World{
        components: make(map[reflect.Type]map[Entity]interface{}),
    }
}

// 제네릭으로 타입 안전하게 추가
func AddComponent[T any](w *World, entity Entity, component T) {
    componentType := reflect.TypeOf(component)

    if w.components[componentType] == nil {
        w.components[componentType] = make(map[Entity]interface{})
    }

    w.components[componentType][entity] = component
}

// 제네릭으로 타입 안전하게 가져오기
func GetComponent[T any](w *World, entity Entity) (T, bool) {
    var zero T
    componentType := reflect.TypeOf(zero)

    if store, exists := w.components[componentType]; exists {
        if comp, found := store[entity]; found {
            return comp.(T), true
        }
    }

    return zero, false
}

// 컴포넌트 존재 여부 확인
func HasComponent[T any](w *World, entity Entity) bool {
    _, exists := GetComponent[T](w, entity)
    return exists
}

// 사용 예시
func main() {
    world := NewWorld()
    player := world.CreateEntity()

    AddComponent(world, player, Position{X: 10, Y: 20})
    AddComponent(world, player, Velocity{DX: 1, DY: 0})

    pos, exists := GetComponent[Position](world, player)
    if exists {
        fmt.Printf("Position: %+v\n", pos)
    }
}
```

**장점**:
- 타입 안전성 유지
- 새 컴포넌트 타입 추가가 쉬움
- 범용적인 코드 작성 가능

**단점**:
- Reflection 사용으로 인한 약간의 성능 오버헤드
- 컴파일 타임 타입 체크 제한적

### 3. Archetype 방식 (Unity DOTS)

같은 컴포넌트 조합을 가진 엔티티들을 함께 묶어서 저장하는 방식입니다.
**최고의 캐시 효율성**을 제공합니다.

```go
// 컴포넌트 조합을 Archetype이라 부름
type Archetype struct {
    signature  ComponentMask     // 어떤 컴포넌트가 있는지
    entities   []Entity          // 이 Archetype에 속한 엔티티들
    positions  []Position        // Position 컴포넌트 배열
    velocities []Velocity        // Velocity 컴포넌트 배열
    entityToIndex map[Entity]int // 엔티티 -> 배열 인덱스
}

type World struct {
    archetypes map[ComponentMask]*Archetype
}

// 예시: Position + Velocity Archetype
func (w *World) GetOrCreateArchetype(mask ComponentMask) *Archetype {
    if arch, exists := w.archetypes[mask]; exists {
        return arch
    }

    arch := &Archetype{
        signature:     mask,
        entities:      []Entity{},
        positions:     []Position{},
        velocities:    []Velocity{},
        entityToIndex: make(map[Entity]int),
    }

    w.archetypes[mask] = arch
    return arch
}

// 엔티티에 컴포넌트 추가 시 적절한 Archetype으로 이동
func (w *World) AddComponent(entity Entity, component interface{}) {
    // 현재 Archetype에서 엔티티 제거
    oldArch := w.findArchetype(entity)
    oldArch.removeEntity(entity)

    // 새로운 컴포넌트 조합으로 Archetype 생성/찾기
    newMask := oldArch.signature | componentToMask(component)
    newArch := w.GetOrCreateArchetype(newMask)

    // 새 Archetype에 엔티티 추가
    newArch.addEntity(entity, component)
}

// 시스템은 Archetype을 직접 순회 (초고속)
func (s *MovementSystem) Update(w *World, dt float64) {
    // Position + Velocity를 가진 Archetype만 순회
    mask := PositionBit | VelocityBit

    for _, arch := range w.archetypes {
        if arch.signature & mask != mask {
            continue // 필요한 컴포넌트가 없으면 스킵
        }

        // 메모리 연속성 최대화! 캐시 히트율 극대화!
        for i := range arch.entities {
            arch.positions[i].X += arch.velocities[i].DX * dt
            arch.positions[i].Y += arch.velocities[i].DY * dt
        }
    }
}
```

**장점**:
- **최고의 캐시 효율성** (컴포넌트가 메모리에 연속 배치)
- 쿼리 성능 최적화
- 병렬 처리에 매우 유리

**단점**:
- 구현이 복잡
- 컴포넌트 추가/제거 시 Archetype 이동 오버헤드
- 메모리 재할당이 빈번할 수 있음

### 4. Bitset + Component 배열 방식

비트마스크로 엔티티가 어떤 컴포넌트를 가졌는지 표시하는 방식입니다.

```go
type ComponentMask uint64

const (
    PositionBit  ComponentMask = 1 << 0  // 0b0001
    VelocityBit  ComponentMask = 1 << 1  // 0b0010
    SpriteBit    ComponentMask = 1 << 2  // 0b0100
    HealthBit    ComponentMask = 1 << 3  // 0b1000
    CollisionBit ComponentMask = 1 << 4  // 0b10000
)

type World struct {
    // 엔티티가 어떤 컴포넌트를 가졌는지 비트로 표시
    componentMasks map[Entity]ComponentMask

    // 컴포넌트별 Dense 배열 (Sparse Set 패턴)
    positions  []Position
    velocities []Velocity

    // 엔티티 ID -> 배열 인덱스 매핑
    positionIndex  map[Entity]int
    velocityIndex  map[Entity]int
}

// 컴포넌트 추가
func (w *World) AddPosition(entity Entity, pos Position) {
    // 배열에 추가
    w.positions = append(w.positions, pos)
    w.positionIndex[entity] = len(w.positions) - 1

    // 비트마스크 업데이트
    w.componentMasks[entity] |= PositionBit
}

// 컴포넌트 확인 (비트 연산으로 초고속)
func (w *World) HasComponent(entity Entity, mask ComponentMask) bool {
    return w.componentMasks[entity] & mask != 0
}

// 쿼리: Position + Velocity를 가진 엔티티 찾기
func (w *World) Query(required ComponentMask) []Entity {
    var result []Entity
    requiredMask := PositionBit | VelocityBit

    for entity, mask := range w.componentMasks {
        // 비트 AND 연산으로 한 번에 확인!
        if mask & requiredMask == requiredMask {
            result = append(result, entity)
        }
    }

    return result
}

// 시스템 예시
func (s *MovementSystem) Update(w *World, dt float64) {
    entities := w.Query(PositionBit | VelocityBit)

    for _, entity := range entities {
        posIdx := w.positionIndex[entity]
        velIdx := w.velocityIndex[entity]

        w.positions[posIdx].X += w.velocities[velIdx].DX * dt
        w.positions[posIdx].Y += w.velocities[velIdx].DY * dt
    }
}
```

**장점**:
- 쿼리가 매우 빠름 (비트 연산)
- 메모리 효율적
- 컴포넌트 존재 여부 확인이 O(1)

**단점**:
- 최대 64개 컴포넌트 타입 제한 (uint64 사용 시)
- 인덱스 관리 복잡
- 컴포넌트 제거 시 배열 재정렬 필요

### 5. 실제 라이브러리 예시: Donburi

Ebitengine 게임 엔진과 잘 통합되는 Donburi 라이브러리의 사용법:

```go
import "github.com/yohamta/donburi"

// 컴포넌트 타입 정의
var (
    Position = donburi.NewComponentType[PositionData]()
    Velocity = donburi.NewComponentType[VelocityData]()
    Sprite   = donburi.NewComponentType[SpriteData]()
)

type PositionData struct {
    X, Y float64
}

type VelocityData struct {
    DX, DY float64
}

type SpriteData struct {
    Image *ebiten.Image
}

func main() {
    // World 생성
    world := donburi.NewWorld()

    // 엔티티 생성 및 컴포넌트 추가
    player := world.Entry(world.Create(Position, Velocity, Sprite))

    // 컴포넌트 데이터 설정
    donburi.SetValue(player, Position, PositionData{X: 100, Y: 100})
    donburi.SetValue(player, Velocity, VelocityData{DX: 2, DY: 0})

    // 컴포넌트 데이터 가져오기
    pos := donburi.Get[PositionData](player, Position)
    fmt.Printf("Position: %+v\n", pos)

    // 쿼리: Position + Velocity를 가진 모든 엔티티
    query := donburi.NewQuery(
        filter.Contains(Position, Velocity),
    )

    // 쿼리 결과 순회
    query.Each(world, func(entry *donburi.Entry) {
        pos := donburi.Get[PositionData](entry, Position)
        vel := donburi.Get[VelocityData](entry, Velocity)

        pos.X += vel.DX
        pos.Y += vel.DY

        donburi.SetValue(entry, Position, *pos)
    })
}

// 시스템 구현
type MovementSystem struct {
    query *donburi.Query
}

func NewMovementSystem() *MovementSystem {
    return &MovementSystem{
        query: donburi.NewQuery(
            filter.Contains(Position, Velocity),
        ),
    }
}

func (s *MovementSystem) Update(world donburi.World, dt float64) {
    s.query.Each(world, func(entry *donburi.Entry) {
        pos := donburi.Get[PositionData](entry, Position)
        vel := donburi.Get[VelocityData](entry, Velocity)

        pos.X += vel.DX * dt
        pos.Y += vel.DY * dt

        donburi.SetValue(entry, Position, *pos)
    })
}
```

### 구현 방식 비교표

| 방식 | 구현 난이도 | 쿼리 성능 | 캐시 효율 | 메모리 효율 | 유연성 |
|------|------------|----------|----------|-----------|--------|
| Component Table | ⭐ 쉬움 | ⭐⭐ 보통 | ⭐⭐ 보통 | ⭐⭐⭐ 좋음 | ⭐⭐ 보통 |
| Go 제네릭 | ⭐⭐ 보통 | ⭐⭐ 보통 | ⭐⭐ 보통 | ⭐⭐⭐ 좋음 | ⭐⭐⭐ 좋음 |
| Archetype | ⭐⭐⭐ 어려움 | ⭐⭐⭐ 매우 빠름 | ⭐⭐⭐ 최고 | ⭐⭐ 보통 | ⭐⭐ 보통 |
| Bitset + 배열 | ⭐⭐⭐ 어려움 | ⭐⭐⭐ 매우 빠름 | ⭐⭐⭐ 좋음 | ⭐⭐⭐ 좋음 | ⭐ 제한적 |
| 라이브러리 | ⭐ 쉬움 | ⭐⭐⭐ 좋음 | ⭐⭐⭐ 좋음 | ⭐⭐⭐ 좋음 | ⭐⭐ 보통 |

### 추천 방식

- **처음 ECS 배우기**: Component Table 방식
- **일반적인 게임**: Go 제네릭 방식 또는 Donburi 라이브러리
- **고성능 필요**: Archetype 방식 또는 arche 라이브러리
- **간단한 프로젝트**: Donburi 라이브러리 (Ebitengine 사용 시)

## 인기있는 Go ECS 라이브러리

- **[arche](https://github.com/mlange-42/arche)**: 고성능 ECS 라이브러리
- **[donburi](https://github.com/yohamta/donburi)**: Ebitengine과 잘 통합되는 ECS
- **[gecs](https://github.com/tutumagi/gecs)**: 간단하고 가벼운 ECS

## 구현 시 주의사항

### 1. 컴포넌트는 데이터만
- 컴포넌트에 메서드나 로직을 넣지 말 것
- 순수한 데이터 구조체로 유지

### 2. 시스템 간 의존성 관리
- 시스템 실행 순서가 중요할 수 있음
- 시스템 간 직접 통신 대신 이벤트 시스템 활용 권장

### 3. 성능 고려사항
- 컴포넌트 쿼리가 빈번하면 캐싱 고려
- 메모리 연속성을 위해 컴포넌트를 배열로 저장

### 4. 과도한 세분화 주의
- 너무 작은 컴포넌트는 오버헤드 증가
- 적절한 크기의 컴포넌트 설계 필요

### 5. 디버깅
- 엔티티 Inspector 도구 구현 권장
- 런타임에 엔티티의 컴포넌트 확인 기능 필요

## 언제 ECS를 사용해야 하나?

### 적합한 경우
- 많은 수의 게임 오브젝트를 다루는 경우
- 오브젝트 타입이 다양하고 자주 변경되는 경우
- 성능이 중요한 경우
- 모듈화와 재사용성이 중요한 경우

### 부적합한 경우
- 매우 간단한 게임 (오버엔지니어링)
- 오브젝트 타입이 고정적이고 적은 경우
- 빠른 프로토타이핑이 필요한 경우

## 참고 자료

- [Unity DOTS (Data-Oriented Technology Stack)](https://unity.com/dots)
- [Bevy Engine (Rust ECS)](https://bevyengine.org/)
- [ECS FAQ](https://github.com/SanderMertens/ecs-faq)
