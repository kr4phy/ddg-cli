# DuckDuckGo 분석 진행도 (2026-04-19)

## 전체 진행 상태

- **요청된 핵심 문서화 작업: 완료**
- **문서 언어 구성: 한/영 병행 완료**
- **추가 실험( `/d.js`, `/t.js` ): 완료**

---

## 완료된 작업

1. **파라미터 사양 정리 완료**
   - `param-spec.md`
   - `param-spec-ko.md`
   - GET(공식 문서 기반) + POST(실제 폼 관찰 기반) 정리

2. **결과 페이지 구조 사양 정리 완료**
   - Lite: `lite-result-spec.md`, `lite-result-spec-ko.md`
   - HTML: `html-result-spec.md`, `html-result-spec-ko.md`
   - Full: `full-result-spec.md`, `full-result-spec-ko.md`

3. **전체 문서 교차 검수/정정 완료**
   - 잘못된 인코딩 설명(Base64 관련) 수정
   - 추측성/권장성 문구 제거 및 기술 사양 중심으로 정리
   - 페이지네이션 및 필드 설명 누락 보완

4. **`/d.js` 전용 API 사양 분리 완료**
   - `djs-api-spec.md`
   - `djs-api-spec-ko.md`
   - 상태 전이, 파라미터 민감도, 페이지네이션 안정성, 로케일 축 비교 반영

5. **`/t.js` 전용 API 사양 분리 완료**
   - `tjs-api-spec.md`
   - `tjs-api-spec-ko.md`
   - `/t.js`↔`/d.js` 호출 순서 및 역할 분리, 응답 특성 반영

---

## 현재 산출물 목록

- `param-spec.md`
- `param-spec-ko.md`
- `lite-result-spec.md`
- `lite-result-spec-ko.md`
- `html-result-spec.md`
- `html-result-spec-ko.md`
- `full-result-spec.md`
- `full-result-spec-ko.md`
- `djs-api-spec.md`
- `djs-api-spec-ko.md`
- `tjs-api-spec.md`
- `tjs-api-spec-ko.md`

---

## 실험/검증 메모

- `/d.js`는 프로브 환경에서 **JavaScript 챌린지 계열 응답**(`isJsaChallenge`, `anomalyDetectionBlock`)이 반복 관찰됨.
- `/t.js`는 프로브 환경에서 **HTTP 200 + 빈 본문**이 관찰됨.
- Playwright MCP는 세션 연결 이슈(DevTools WebSocket handshake)로 런타임 DOM 기반 검증이 제한되어, 문서는 **HTTP 응답/부트스트랩 스크립트 관찰 근거 중심**으로 작성됨.
