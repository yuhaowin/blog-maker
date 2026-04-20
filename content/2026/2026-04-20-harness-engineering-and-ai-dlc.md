# Harness Engineering 与 AI-DLC：从代码执行到项目迭代

> 同一个模型，不同的 Harness，SWE-bench 上差了 17 题。当模型能力趋于同质化，真正决定 AI Agent 表现上限的，不是它有多聪明，而是它被放在什么样的环境里工作。

## 一、为什么 Harness 比模型重要

2026 年，AI 工程领域达成了一个共识：**Harness 比模型重要**。

### 1.1 同一个模型，差了 17 题

2026 年初，一组 SWE-bench Verified 的评测数据引起了广泛讨论：Augment、Cursor、Claude Code 三个产品，都跑的 Claude Opus
4.5，731 道题，成绩差了 17 题。

模型完全一样。差异来自哪里？**Harness**——包裹在模型外面的那层系统：工具定义、上下文管理、错误恢复、验证循环、子任务编排。

这不是个例。整个行业都在讲同一个故事：模型是 CPU，Harness 是操作系统。没有操作系统，CPU 再快也只是一块芯片。

### 1.2 从 Prompt 到 Context 到 Harness：三次范式转移

![AI 工程演进的三次范式转移](https://image.yuhaowin.com/2026/04/20/231151.png)

| 阶段                  | 时间        | 核心问题                        |
|---------------------|-----------|-----------------------------|
| Prompt Engineering  | 2022–2024 | 怎么写好一条指令                    |
| Context Engineering | 2025      | 怎么策展所有相关信息（RAG、Memory、工具描述） |
| Harness Engineering | 2026      | 怎么设计环境、约束和反馈循环              |

Mitchell Hashimoto（HashiCorp 创始人）在 2026 年 2 月首次明确了"Harness Engineering"这个概念：

> "每次 agent 犯错，不要寄希望于'下次做对'。改造环境，让它不可能再用同样的方式犯错。"

这句话精准地定义了 harness 的本质：**不是教 agent 做什么，而是让环境保证 agent 只能做对的事**。

### 1.3 行业共识的形成

2026 年 2 月起，这个认知从个人观点变成了行业共识：

- **OpenAI** 用 Codex agent 从空 repo 构建完整产品，发现 harness 的工程设计决定了 agent 能否长时间可靠运行
- **Birgitta Böckeler** 在 Martin Fowler 的网站上撰文，将 Harness Engineering 定位为软件工程实践的新分支
- **Anthropic** 提出 GAN 式 Generator/Evaluator 架构，核心发现是模型不能可靠地评估自己的工作，必须由 harness 提供外部验证环
- **Stripe** 的工程师总结——"The Walls Matter More Than the Model"，围墙比引擎重要

> "2025 Was Agents. 2026 Is Agent Harnesses." — Aakash Gupta

---

## 二、代码级 Harness：以 Claude Code 为例

当前主流的 agent harness 都聚焦在**代码执行层**。Claude Code 构建了最完整的六层架构：

![Claude Code 六层 Harness 架构](https://image.yuhaowin.com/2026/04/20/231210.png)

1. **CLAUDE.md** — 项目上下文定义
2. **Tools/MCP** — 能力接入层
3. **Skills** — 方法论封装
4. **Hooks** — 机械约束
5. **Subagents** — 隔离工作者
6. **Verifiers** — 验证循环

**Codex** 走的是云沙箱路线：Agent 拿到一个空白环境，读代码、做计划、写代码、跑测试、交 PR。GPT-5.3-Codex 跑了 25 小时不间断，13M
token，30K 行代码。

**Cursor** 是 IDE 原生集成：实时协作，视觉反馈，360K 付费用户。

这些 harness 各有所长，但解决的问题是同一类：**agent 怎么写代码**。包括怎么读文件、怎么调工具、怎么跑测试、出错了怎么恢复、上下文满了怎么压缩。

---

## 三、为什么需要项目管理 Harness

当 agent 从单任务（修一个 bug）升级到多任务（做一个 feature），从单 agent 升级到多 agent 团队协作，代码级 harness 就不够用了。

### 缺失的环节

- **需求理解**：这个任务从哪来？需求是否被充分理解？agent 是在正确的理解上执行，还是在错误的假设上高效产出垃圾？
- **任务编排**：5 个 agent 同时工作时，谁干什么？依赖关系是什么？两个 agent 同时抢一个任务怎么办？
- **验收闭环**：任务完成后，谁来验证？验证标准是什么？agent 自己说"做完了"可信吗？
- **迭代节奏**：一轮做完后，下一轮自动开始了吗？下游任务知道上游已完成了吗？

类比一下：现有的 harness 给了 agent 一个配置齐全的**工位**——双屏显示器、机械键盘、IDE 全装好。但没有给它一个**项目部**
——没有需求评审、没有任务看板、没有 Sprint 节奏、没有验收标准。

Agent 知道怎么敲键盘，但不知道为什么敲、敲完给谁看、下一步做什么。

---

## 四、项目管理 Harness：AI-DLC 的核心理论

AWS 提出的 **AI-Driven Development Lifecycle (AI-DLC)** 为项目管理 Harness 提供了系统性的理论框架。AI-DLC 不是简单地将 AI
嫁接到传统敏捷方法上，而是**重新想象**整个开发生命周期。

### 4.1 八大核心原则

![AI-DLC 八大核心原则](https://image.yuhaowin.com/2026/04/20/231227.png)

1. **Reimagine Rather Than Retrofit（重新想象而非改造）**
    - 传统方法为长周期（数周数月）设计， rituals 如每日站会、回顾会
    - AI-DLC 的迭代以小时或天计算，需要实时验证和反馈机制
    - 不再纠结故事点估算，而是关注业务价值交付

2. **Reverse the Conversation Direction（反转对话方向）**
    - 传统：人类写 Prompt → AI 执行 → 人类检查结果 → 不满意改 Prompt → AI 重来
    - AI-DLC：**AI 提问 → 人类回答 → AI 验证自洽性 → 有矛盾就追问 → 共识 → 再开干**
    - 人类从"执行者"变为"审批者"，从"写需求"变为"回答 AI 的澄清问题"

3. **Domain-Driven by Default（默认领域驱动）**
    - AI 自动应用 DDD 原则，将系统分解为独立、适度大小的限界上下文
    - 开发者只需验证和调整，无需手动进行领域建模

4. **Align with AI Capability（与 AI 能力对齐）**
    - 认识到当前 AI 还不能完全自主地将高层意图转化为可执行代码
    - 采用"AI-Driven"范式而非"AI-Assisted"：AI 主动编排，人类负责验证和决策

5. **Cater to Building Complex Systems（服务于复杂系统构建）**
    - 针对需要持续功能适配、高架构复杂性、多权衡管理的系统
    - 涉及多个团队在大规模/受监管组织内的协作

6. **Retain What Enhances Human Symbiosis（保留增强人机协作的元素）**
    - 保留用户故事（作为人类与 AI 理解的契约）
    - 保留风险登记册（确保 AI 生成的计划和代码符合组织风险框架）
    - 将这些元素优化为实时使用

7. **Facilitate Transition Through Familiarity（通过熟悉度促进过渡）**
    - 任何现有从业者都应该能在一天内上手
    - 保留熟悉术语的底层关系，同时引入现代化术语
    - 例如：将 Scrum 中的"Sprint"重新命名为"Bolt"，强调快速、集中的交付节奏

8. **Streamline Responsibilities for Efficiency（简化职责以提高效率）**
    - 利用 AI 的任务分解和决策能力，开发者超越传统专业壁垒（前端、后端、DevOps、安全）
    - 减少多专业角色的需求，简化开发流程

### 4.2 核心工件（Artefacts）

| 工件                  | 定义                                                         |
|---------------------|------------------------------------------------------------|
| **Intent**          | 高层目的陈述，封装需要实现的目标（业务目标、功能或技术成果）                             |
| **Unit**            | 从 Intent 分解出的内聚、自包含工作单元，可独立开发和部署（类似 DDD 的子域或 Scrum 的 Epic） |
| **Bolt**            | AI-DLC 中最小的迭代单元，以小时或天为单位的构建-验证周期（类比 Sprint，但更短更快）          |
| **Domain Design**   | 使用 DDD 原则建模的核心业务逻辑，包括聚合、值对象、实体、领域事件等                       |
| **Logical Design**  | 扩展 Domain Design 以满足非功能需求，应用架构设计模式（CQRS、熔断器等）              |
| **Deployment Unit** | operational 工件，包含打包的可执行代码、配置和基础设施组件，经过功能验收、安全、NFR 和风险的严格测试 |

### 4.3 三个阶段与仪式

**Inception Phase（起始阶段）**

- 捕获 Intent 并转化为 Units
- **Mob Elaboration（群体细化）仪式**：Product Owner、开发者、QA 等利益相关者一起审查和细化 AI 生成的工件
- AI 提出澄清问题，团队验证答案的自洽性
- 输出：明确定义的 Units 及其组件（PRFAQ、用户故事、NFR、风险描述、度量标准、建议的 Bolts）

**Construction Phase（构建阶段）**

- 将 Units 迭代执行为可测试、可运维的 Deployment Units
- **Mob Construction（群体构建）仪式**：所有团队共置一室，类似于 Mob Elaboration
- 进展：Domain Design → Logical Design → Code Generation → Automated Testing
- 开发者专注于验证 AI 生成的输出并做出关键决策

**Operations Phase（运维阶段）**

- 部署、可观测性和系统维护
- AI 分析遥测数据（指标、日志、追踪）以检测模式、识别异常、预测潜在 SLA 违规
- AI 与预定义的事件手册集成，提出可操作的推荐（如资源扩展、性能调优、故障隔离）

---

## 五、Chorus：AI-DLC 的项目管理 Harness 实践

Chorus（https://chorus-ai.dev）是一个开源的 AI-DLC Agent 协作平台，它在 Claude Code 等代码级 harness 之上，提供**项目级
harness**，让 agent 拥有从想法到验收的完整迭代环。

### 5.1 六阶段结构化流水线

Chorus 实现了 AI-DLC 的完整管道：

![Chorus AI-DLC 六阶段流水线](https://image.yuhaowin.com/2026/04/20/231240.png)

| 阶段              | 谁在做             | 做什么                                                           |
|-----------------|-----------------|---------------------------------------------------------------|
| **Idea**        | 人类              | 抛出一个想法，可以很粗糙                                                  |
| **Elaboration** | PM Agent → 人类   | AI 不直接开干，而是向人类提问："目标用户规模？""需要离线支持吗？"人类回答，AI 验证自洽性，有矛盾就追问，直到共识 |
| **Proposal**    | PM Agent        | 产出 PRD 文档草案 + 任务 DAG（依赖图）                                     |
| **Approval**    | Admin / 人类      | 审批方案，通过后任务才实体化                                                |
| **Execute**     | Developer Agent | 认领任务，在 Claude Code 中执行，自检验收标准后提交                              |
| **Verify**      | Admin / 人类      | 逐条验证验收标准，通过或打回。下游任务自动 unblock，下一波开始                           |

这不是一个"任务管理看板"。这是一个**让 agent 知道自己在整个项目中处于什么位置的运行时环境**。

每个阶段的边界都是 harness 级别的约束：

- 需求没细化完，**开不了工**
- 方案没审批，**任务不存在**
- 上游任务没验收，**下游任务不会 unblock**
- 做完没过验收，**不算 Done**

这就是 Hashimoto 说的那件事：不是教 agent "你应该先理解需求再动手"——**环境保证了它必须先理解需求才能动手**。

### 5.2 Reversed Conversation（反转对话）

传统工作流的信息流向是单向的：

```
人写 Prompt → AI 执行 → 人检查结果 → 不满意改 Prompt → AI 重来
```

这个模式的致命问题：**agent 在错误的理解上高效执行**。它可能写了 500 行完美的代码，但解决的是错误的问题。

Chorus 的 Elaboration 机制反转了对话方向：

```
人提想法 → AI 提问 → 人回答 → AI 验证答案自洽性 → 有矛盾就追问 → 共识 → 再开干
```

PM Agent 读完一个 Idea 后，不是直接开干，而是生成一组结构化问题。如果人回答了"需要离线支持"但又说"要实时同步"，PM
会追问——因为这两个需求在某些场景下是矛盾的。

> Harness 的价值不只是"让 agent 做得快"，更是"让 agent 做对的事"。Elaboration 是 Chorus 在 harness 层面对需求质量的保障：不是靠
> agent 的"理解力"，而是靠**结构化问答的流程约束**。

### 5.3 DAG + Wave 验证：多 Agent 并行不乱序

当一个 Proposal 产出 8 个任务、3 层依赖时，Chorus 构建 Task DAG（有向无环图），并用 Wave 模型管理执行节奏：

![DAG + Wave 任务依赖管理](https://image.yuhaowin.com/2026/04/20/231252.png)

```
Wave 1: [Task A] [Task B] [Task C]  ← 无依赖，可并行
         ↓         ↓
Wave 2:      [Task D] [Task E]      ← 依赖 Wave 1 的任务
                  ↓
Wave 3:          [Task F]           ← 依赖 Task E
```

关键设计决策：**不是在执行时强制阻塞，而是在验证时卡住——上游没验收，下游就不会开放**。

- Wave 1 的任务可以被多个 agent 并行认领执行
- 每个 agent 完成后提交验收
- Wave 1 全部验收通过后，Wave 2 自动 unblock
- 如果 Wave 1 某个任务验收失败被打回，依赖它的下游任务不会 unblock

这正是 Stripe 说的 "The Walls Matter More Than the Model"：DAG 就是墙。Agent 不需要"理解"依赖关系——**环境本身阻止了乱序执行
**。

### 5.4 三种专业 Agent 角色

Chorus 定义了三种专业化 agent 角色，对应 AI-DLC 的不同阶段：

| 角色                  | 职责                                  | 对应 AI-DLC 阶段       |
|---------------------|-------------------------------------|--------------------|
| **PM Agent**        | 分析想法、编写 PRD、设计任务 DAG、进行 Elaboration | Inception Phase    |
| **Developer Agent** | 认领任务、执行代码、多子 agent 群集模式工作           | Construction Phase |
| **Admin Agent**     | 创建项目、审批方案、进行最终验收审核                  | Approval & Verify  |

### 5.5 验收不是可选项

Anthropic 的工程博客指出：**模型不能可靠地评估自己的工作**。

Chorus 在项目层面实现了 GAN 式的 Generator/Evaluator 架构：

1. **Developer Agent 完成任务后，先跑 Acceptance Criteria 自检**——逐条对照验收标准，标记每一条是否满足
2. **自检通过后提交验收，由 Admin 或人类逐条确认**——不是 agent 自己说了算
3. **验证失败可以打回**——附带反馈，agent 修改后重新提交

> Agent 说"做完了"，和 Admin 验证过"确实做完了"，是两件完全不同的事。Chorus 把这个区分编码成了 harness
> 的一部分，不依赖任何人"记得去检查"。

### 5.6 核心优势总结

| 优势                | 说明                                           |
|-------------------|----------------------------------------------|
| **零上下文注入**        | 无需手动 prompt engineering，agent 自动获取项目上下文      |
| **多 Agent 可观测性**  | 实时追踪 session 级别的活动，了解谁在做什么                   |
| **MCP-native 架构** | 支持 Bring Your Own Agent，兼容任何遵循 MCP 协议的 agent |
| **AI-DLC 完整实现**   | 从需求澄清到任务验收的结构化流水线，交付项目而不只是写代码                |
| **"AI 提议，人类把关"**  | 反转对话范式，人类从执行者变为审批者                           |

---

## 六、结语

> "2025 Was Agents. 2026 Is Agent Harnesses." — Aakash Gupta

这句话需要一个补充：

2026 年的 Harness Engineering 有两层。**第一层是代码级 harness**——Claude Code、Codex、Cursor 已经做得很好。**第二层是项目级
harness**——从想法细化到任务验收的完整迭代环境——这是正在被填补的空白。

| 层次          | 解决的问题       | 代表                         |
|-------------|-------------|----------------------------|
| 代码级 Harness | Agent 怎么写代码 | Claude Code, Codex, Cursor |
| 项目级 Harness | Agent 怎么做项目 | Chorus                     |

两层结合，agent 才拥有完整的工作环境：知道做什么（Idea + Elaboration）、怎么做（Code Harness）、做完给谁看（Verify）、下一步是什么（DAG
unblock）。

当模型能力越来越强、越来越同质化，决定 agent 上限的不再是它有多聪明，而是它被放在什么样的环境里工作。

**Harness 不是辅助。Harness 是上限。**

---

## 参考与引用

- [Chorus — AI-DLC Agent Collaboration Platform](https://github.com/Chorus-AIDLC/Chorus)
- [Chorus 官方文档](https://chorus-ai.dev/zh/)
- AWS, "AI-Driven Development Lifecycle (AI-DLC) Method Definition"
- Mitchell Hashimoto, "My AI Adoption Journey — Step 5: Engineer the Harness"
- OpenAI, "Harness engineering: leveraging Codex in an agent-first world"
- Birgitta Böckeler / Martin Fowler, "Harness Engineering"
- Anthropic Engineering, "Effective harnesses for long-running agents"
- Anup Jadhav, "Stripe's coding agents: the walls matter more than the model"
- Aakash Gupta, "2025 Was Agents. 2026 Is Agent Harnesses."
