# glfw.WindowHint

## 1. 윈도우 속성 관련

### glfw.Resizable
- **기본값**: `glfw.True`
- **설명**: 윈도우 크기 조절 가능 여부

### glfw.Visible
- **기본값**: `glfw.True`
- **설명**: 윈도우 생성 시 표시 여부

### glfw.Decorated
- **기본값**: `glfw.True`
- **설명**: 윈도우 테두리 및 타이틀바 표시 여부

### glfw.Focused
- **기본값**: `glfw.True`
- **설명**: 윈도우 생성 시 포커스 획득 여부

### glfw.AutoIconify
- **기본값**: `glfw.True`
- **설명**: 전체화면 윈도우가 포커스 잃을 때 최소화 여부

### glfw.Floating
- **기본값**: `glfw.False`
- **설명**: 윈도우를 항상 최상위에 표시 여부

### glfw.Maximized
- **기본값**: `glfw.False`
- **설명**: 윈도우 생성 시 최대화 상태

### glfw.CenterCursor
- **기본값**: `glfw.True`
- **설명**: 전체화면 모드에서 커서 중앙 배치

### glfw.TransparentFramebuffer
- **기본값**: `glfw.False`
- **설명**: 프레임버퍼 투명도 지원 여부

### glfw.FocusOnShow
- **기본값**: `glfw.True`
- **설명**: Show() 호출 시 포커스 획득 여부

### glfw.ScaleToMonitor
- **기본값**: `glfw.False`
- **설명**: 콘텐츠 스케일을 모니터 DPI에 맞춤

## 2. 프레임버퍼 관련

### glfw.RedBits
- **기본값**: `8`
- **설명**: 빨강 채널 비트 수

### glfw.GreenBits
- **기본값**: `8`
- **설명**: 녹색 채널 비트 수

### glfw.BlueBits
- **기본값**: `8`
- **설명**: 파랑 채널 비트 수

### glfw.AlphaBits
- **기본값**: `8`
- **설명**: 알파 채널 비트 수

### glfw.DepthBits
- **기본값**: `24`
- **설명**: 깊이 버퍼 비트 수

### glfw.StencilBits
- **기본값**: `8`
- **설명**: 스텐실 버퍼 비트 수

### glfw.Samples
- **기본값**: `0`
- **설명**: MSAA 샘플 수 (안티앨리어싱, 0은 비활성화)

### glfw.SRGBCapable
- **기본값**: `glfw.False`
- **설명**: sRGB 색공간 지원 여부

### glfw.DoubleBuffer
- **기본값**: `glfw.True`
- **설명**: 더블 버퍼링 사용 여부

## 3. OpenGL 컨텍스트 관련

### glfw.ContextVersionMajor
- **기본값**: `1`
- **설명**: OpenGL 메이저 버전 (예: 4)

### glfw.ContextVersionMinor
- **기본값**: `0`
- **설명**: OpenGL 마이너 버전 (예: 6)

### glfw.OpenGLProfile
- **기본값**: `glfw.OpenGLAnyProfile`
- **설명**: OpenGL 프로파일 선택
  - `glfw.OpenGLAnyProfile`: 어떤 프로파일이든 가능
  - `glfw.OpenGLCoreProfile`: 코어 프로파일 (레거시 기능 제거)
  - `glfw.OpenGLCompatProfile`: 호환성 프로파일 (레거시 포함)

### glfw.OpenGLForwardCompatible
- **기본값**: `glfw.False`
- **설명**: 하위 호환성 제거 (macOS에서 OpenGL 3.2+ 사용 시 필수)

### glfw.OpenGLDebugContext
- **기본값**: `glfw.False`
- **설명**: 디버그 컨텍스트 생성 여부

### glfw.ContextRobustness
- **기본값**: `glfw.NoRobustness`
- **설명**: 컨텍스트 견고성 레벨
  - `glfw.NoRobustness`: 견고성 없음
  - `glfw.NoResetNotification`: 리셋 알림 없음
  - `glfw.LoseContextOnReset`: 리셋 시 컨텍스트 손실

### glfw.ContextReleaseBehavior
- **기본값**: `glfw.AnyReleaseBehavior`
- **설명**: 컨텍스트 해제 동작 방식

### glfw.ContextNoError
- **기본값**: `glfw.False`
- **설명**: 에러 생성 비활성화 (성능 향상, 디버깅 불가)

### glfw.ContextCreationAPI
- **기본값**: `glfw.NativeContextAPI`
- **설명**: 컨텍스트 생성 API
  - `glfw.NativeContextAPI`: 네이티브 API
  - `glfw.EGLContextAPI`: EGL API
  - `glfw.OSMesaContextAPI`: OSMesa API

## 4. 모니터/리프레시 관련

### glfw.RefreshRate
- **기본값**: `glfw.DontCare`
- **설명**: 전체화면 모드 리프레시율 (Hz)

## 주의사항

- WindowHint는 반드시 `glfw.CreateWindow()` **이전**에 호출
- 잘못된 조합은 윈도우 생성 실패를 일으킬 수 있습니다
- macOS에서 OpenGL 3.2+ 사용 시 `OpenGLForwardCompatible`를 True로 설정 필요
- 모든 힌트가 모든 플랫폼에서 지원되는 것은 아닙니다
