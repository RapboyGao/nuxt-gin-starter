/**
 * =====================================================
 * Nuxt Gin WebSocket Client
 * -----------------------------------------------------
 * This file is auto-generated. Do not edit by hand.
 * Regenerate by running the Go server endpoint export.
 * Edits will be overwritten on the next generation.
 * -----------------------------------------------------
 * 本文件由工具自动生成，请勿手动修改。
 * 如需更新，请通过 Go 服务端重新生成。
 * 手动修改将在下次生成时被覆盖。
 * =====================================================
 */

const isPlainObject = (value: unknown): value is Record<string, unknown> =>
  Object.prototype.toString.call(value) === '[object Object]';

const isoDateLike =
  /^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(?:\.\d{1,9})?(?:Z|[+\-]\d{2}:\d{2})$/;

const normalizeWsRequestJSON = (value: unknown): unknown => {
  if (value instanceof Date) return value.toISOString();
  if (Array.isArray(value)) return value.map(normalizeWsRequestJSON);
  if (isPlainObject(value)) {
    const out: Record<string, unknown> = {};
    for (const [k, v] of Object.entries(value))
      out[k] = normalizeWsRequestJSON(v);
    return out;
  }
  return value;
};

const normalizeWsResponseJSON = (value: unknown): unknown => {
  if (Array.isArray(value)) return value.map(normalizeWsResponseJSON);
  if (typeof value === 'string' && isoDateLike.test(value)) {
    const date = new Date(value);
    if (!Number.isNaN(date.getTime())) return date;
  }
  if (isPlainObject(value)) {
    const out: Record<string, unknown> = {};
    for (const [k, v] of Object.entries(value))
      out[k] = normalizeWsResponseJSON(v);
    return out;
  }
  return value;
};

export interface WebSocketConvertOptions<TSend = unknown, TReceive = unknown> {
  serialize?: (value: TSend) => unknown;
  deserialize?: (value: unknown) => TReceive;
}

export interface TypedHandlerOptions<TReceive, TPayload> {
  selectPayload?: (message: TReceive) => unknown;
  decode?: (payload: unknown) => TPayload;
  validate?: (payload: unknown, message: TReceive) => boolean;
}

export interface TypeHandlerOptions<TReceive> {
  validate?: (message: TReceive) => boolean;
}

const resolveWebSocketURL = (url: string): string => {
  if (url.startsWith('ws://') || url.startsWith('wss://')) return url;
  if (url.startsWith('http://')) return `ws://${url.slice(7)}`;
  if (url.startsWith('https://')) return `wss://${url.slice(8)}`;
  if (url.startsWith('/')) {
    const isHttps =
      typeof window !== 'undefined' && window.location?.protocol === 'https:';
    const host = typeof window !== 'undefined' ? window.location.host : '';
    return `${isHttps ? 'wss' : 'ws'}://${host}${url}`;
  }
  return url;
};

const joinURLPath = (baseURL: string, path: string): string => {
  const base = baseURL.trim();
  const p = path.trim();
  if (!base) return p.startsWith('/') ? p : `/${p}`;
  if (!p)
    return base.startsWith('/')
      ? base.replace(/\/+$/, '')
      : `/${base.replace(/\/+$/, '')}`;
  const trimmedBase = base.replace(/\/+$/, '');
  const trimmedPath = p.replace(/^\/+/, '');
  return trimmedBase.startsWith('/')
    ? `${trimmedBase}/${trimmedPath}`
    : `/${trimmedBase}/${trimmedPath}`;
};

export class TypedWebSocketClient<
  TReceive = unknown,
  TSend = unknown,
  TType extends string = string,
> {
  public readonly socket: WebSocket;
  private readonly serialize: (value: TSend) => unknown;
  private readonly deserialize: (value: unknown) => TReceive;
  private readonly messageListeners = new Set<(message: TReceive) => void>();
  private readonly openListeners = new Set<(event: Event) => void>();
  private readonly closeListeners = new Set<(event: CloseEvent) => void>();
  private readonly errorListeners = new Set<(event: Event) => void>();
  private readonly typedListeners = new Map<
    TType,
    Set<(message: TReceive) => void>
  >();

  constructor(url: string, options: WebSocketConvertOptions<TSend, TReceive>) {
    this.socket = new WebSocket(resolveWebSocketURL(url));
    this.serialize =
      options?.serialize ?? ((value: TSend) => normalizeWsRequestJSON(value));
    this.deserialize =
      options?.deserialize ??
      ((value: unknown) => normalizeWsResponseJSON(value) as TReceive);

    this.socket.addEventListener('message', (event) => {
      let payload: unknown = event.data;
      if (typeof payload === 'string') {
        try {
          payload = JSON.parse(payload);
        } catch {
          // keep raw payload
        }
      }
      const message = this.deserialize(payload);
      this.emitMessage(message);
    });
    this.socket.addEventListener('open', (event) => {
      for (const listener of this.openListeners) listener(event);
    });
    this.socket.addEventListener('close', (event) => {
      for (const listener of this.closeListeners) listener(event);
    });
    this.socket.addEventListener('error', (event) => {
      for (const listener of this.errorListeners) listener(event);
    });
  }

  send(message: TSend): void {
    const data = this.serialize(message);
    this.socket.send(JSON.stringify(data));
  }

  close(): void {
    this.socket.close();
  }

  onMessage(handler: (message: TReceive) => void): () => void {
    this.messageListeners.add(handler);
    return () => this.messageListeners.delete(handler);
  }

  onOpen(handler: (event: Event) => void): () => void {
    this.openListeners.add(handler);
    return () => this.openListeners.delete(handler);
  }

  onClose(handler: (event: CloseEvent) => void): () => void {
    this.closeListeners.add(handler);
    return () => this.closeListeners.delete(handler);
  }

  onError(handler: (event: Event) => void): () => void {
    this.errorListeners.add(handler);
    return () => this.errorListeners.delete(handler);
  }

  onType(
    type: TType,
    handler: (message: TReceive) => void,
    options?: TypeHandlerOptions<TReceive>
  ): () => void {
    const listeners =
      this.typedListeners.get(type) ?? new Set<(message: TReceive) => void>();
    const wrapped = (message: TReceive) => {
      if (options?.validate && !options.validate(message)) return;
      handler(message);
    };
    listeners.add(wrapped);
    this.typedListeners.set(type, listeners);
    return () => {
      const current = this.typedListeners.get(type);
      if (!current) return;
      current.delete(wrapped);
      if (current.size === 0) this.typedListeners.delete(type);
    };
  }

  onTyped<TPayload>(
    type: TType,
    handler: (payload: TPayload, message: TReceive) => void,
    options?: TypedHandlerOptions<TReceive, TPayload>
  ): () => void {
    return this.onType(type, (message) => {
      const rawPayload = options?.selectPayload
        ? options.selectPayload(message)
        : this.defaultPayload(message);
      if (options?.validate && !options.validate(rawPayload, message)) return;
      const payload = options?.decode
        ? options.decode(rawPayload)
        : (rawPayload as TPayload);
      handler(payload, message);
    });
  }

  private emitMessage(message: TReceive): void {
    for (const listener of this.messageListeners) {
      try {
        listener(message);
      } catch {
        // ignore single listener errors and continue dispatch
      }
    }
    const type = this.defaultMessageType(message);
    if (!type) return;
    const listeners = this.typedListeners.get(type);
    if (!listeners) return;
    for (const listener of listeners) {
      try {
        listener(message);
      } catch {
        // ignore single listener errors and continue dispatch
      }
    }
  }

  private defaultMessageType(message: TReceive): TType | undefined {
    if (!isPlainObject(message)) return undefined;
    const value = (message as Record<string, unknown>)['type'];
    return typeof value === 'string' ? (value as TType) : undefined;
  }

  private defaultPayload(message: TReceive): unknown {
    if (!isPlainObject(message)) return message;
    return (message as Record<string, unknown>)['payload'];
  }
}

/**
 * WebSocket demo with typed message handlers
 */
// Literal union is emitted as type because interface cannot model union values.
// 字面量联合类型使用 type，因为 interface 不能表达联合值。
export type ChatDemoMessageType = string;
export function chatDemo<TSend = WebSocketMessage>(
  options: WebSocketConvertOptions<TSend, WsServerEnvelope>
): TypedWebSocketClient<WsServerEnvelope, TSend, ChatDemoMessageType> {
  const url = joinURLPath('/ws-go/v1', '/chat-demo');
  return new TypedWebSocketClient<WsServerEnvelope, TSend, ChatDemoMessageType>(
    url,
    options
  );
}

// =====================================================
// INTERFACES & VALIDATORS
// Default: object schemas use interface.
// Fallback: use type only when interface cannot model the shape.
// 默认：对象结构使用 interface。
// 兜底：只有 interface 无法表达时才使用 type。
// =====================================================

// -----------------------------------------------------
// TYPE: WebSocketMessage
// -----------------------------------------------------
export interface WebSocketMessage {
  type: string;
  payload: string;
}

/**
 * Validate whether a value matches WebSocketMessage.
 * 校验一个值是否符合 WebSocketMessage 结构。
 */
export function validateWebSocketMessage(
  value: unknown
): value is WebSocketMessage {
  if (!isPlainObject(value)) return false;
  const obj = value as Record<string, unknown>;
  if (!('type' in obj)) return false;
  if (!(typeof obj['type'] === 'string')) return false;
  if (!('payload' in obj)) return false;
  if (!(typeof obj['payload'] === 'string')) return false;
  return true;
}

/**
 * Ensure a typed WebSocketMessage after validation.
 * 先校验，再确保得到类型化的 WebSocketMessage。
 */
export function ensureWebSocketMessage(value: unknown): WebSocketMessage {
  if (!validateWebSocketMessage(value)) {
    throw new Error('Invalid WebSocketMessage');
  }
  return value;
}

// -----------------------------------------------------
// TYPE: WsServerEnvelope
// -----------------------------------------------------
export interface WsServerEnvelope {
  type: string;
  client: string;
  message: string;
  at: string;
}

/**
 * Validate whether a value matches WsServerEnvelope.
 * 校验一个值是否符合 WsServerEnvelope 结构。
 */
export function validateWsServerEnvelope(
  value: unknown
): value is WsServerEnvelope {
  if (!isPlainObject(value)) return false;
  const obj = value as Record<string, unknown>;
  if (!('type' in obj)) return false;
  if (!(typeof obj['type'] === 'string')) return false;
  if (!('client' in obj)) return false;
  if (!(typeof obj['client'] === 'string')) return false;
  if (!('message' in obj)) return false;
  if (!(typeof obj['message'] === 'string')) return false;
  if (!('at' in obj)) return false;
  if (!(typeof obj['at'] === 'string')) return false;
  return true;
}

/**
 * Ensure a typed WsServerEnvelope after validation.
 * 先校验，再确保得到类型化的 WsServerEnvelope。
 */
export function ensureWsServerEnvelope(value: unknown): WsServerEnvelope {
  if (!validateWsServerEnvelope(value)) {
    throw new Error('Invalid WsServerEnvelope');
  }
  return value;
}
