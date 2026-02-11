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

import type { ProductCreateRequest, ProductListQueryParams, WebSocketMessage, WsProductDeletePayload, WsProductUpdatePayload } from './auto-generated-types';


// #region Runtime Helpers
// =====================================================

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

const isDevelopmentEnv = (): boolean => {
  if (typeof import.meta !== 'undefined' && (import.meta as any)?.env) {
    const dev = (import.meta as any).env?.DEV;
    if (typeof dev === 'boolean') return dev;
  }
  return false;
};

const resolveGinPort = (): string => {
  if (typeof window !== 'undefined') {
    const ginPort = useRuntimeConfig().public.ginPort;
    if (
      ginPort !== undefined &&
      ginPort !== null &&
      String(ginPort).trim() !== ''
    ) {
      return String(ginPort);
    }
    if (window.location?.port && window.location.port.trim() !== '') {
      return window.location.port;
    }
    return window.location?.protocol === 'https:' ? '443' : '80';
  }
  if (
    typeof import.meta !== 'undefined' &&
    (import.meta as any)?.env?.NUXT_GIN_PORT
  ) {
    return String((import.meta as any).env.NUXT_GIN_PORT);
  }
  return '80';
};

const resolveWebSocketURL = (url: string): string => {
  if (url.startsWith('ws://') || url.startsWith('wss://')) return url;
  if (url.startsWith('http://')) return `ws://${url.slice(7)}`;
  if (url.startsWith('https://')) return `wss://${url.slice(8)}`;
  if (url.startsWith('/')) {
    const isHttps =
      typeof window !== 'undefined' && window.location?.protocol === 'https:';
    const protocol = isHttps ? 'wss' : 'ws';
    if (typeof window !== 'undefined') {
      if (isDevelopmentEnv()) {
        return `${protocol}://${window.location.hostname}:${resolveGinPort()}${url}`;
      }
      return `${protocol}://${window.location.host}${url}`;
    }
    return url;
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

// #endregion Runtime Helpers

// #region Typed WebSocket Client
// =====================================================

/**
 * Generic typed WebSocket client with message and type-based subscriptions.
 * 通用的类型化 WebSocket 客户端，支持全量消息订阅与按 type 订阅。
 */
export class TypedWebSocketClient<
  TReceive = unknown,
  TSend = unknown,
  TType extends string = string,
> {
  public readonly socket: WebSocket;
  public readonly url: string;
  public status: 'connecting' | 'open' | 'closing' | 'closed' = 'connecting';
  public lastError?: Event;
  public lastClose?: CloseEvent;
  public connectedAt?: Date;
  public closedAt?: Date;
  public messagesSent = 0;
  public messagesReceived = 0;
  public reconnectCount = 0;
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

  /**
   * Create a websocket client and connect immediately.
   * 创建 websocket 客户端并立即发起连接。
   */
  constructor(url: string, options: WebSocketConvertOptions<TSend, TReceive>) {
    const resolvedURL = resolveWebSocketURL(url);
    this.url = resolvedURL;
    this.socket = new WebSocket(resolvedURL);
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
      this.messagesReceived += 1;
      this.emitMessage(message);
    });
    this.socket.addEventListener('open', (event) => {
      this.status = 'open';
      this.connectedAt = new Date();
      this.closedAt = undefined;
      for (const listener of this.openListeners) listener(event);
    });
    this.socket.addEventListener('close', (event) => {
      this.status = 'closed';
      this.lastClose = event;
      this.closedAt = new Date();
      for (const listener of this.closeListeners) listener(event);
    });
    this.socket.addEventListener('error', (event) => {
      this.lastError = event;
      for (const listener of this.errorListeners) listener(event);
    });
  }

  /**
   * Current WebSocket readyState.
   * 当前 WebSocket 连接状态。
   */
  get readyState(): number {
    return this.socket.readyState;
  }

  /**
   * Whether the socket is currently open.
   * 当前连接是否处于打开状态。
   */
  get isOpen(): boolean {
    return this.readyState === WebSocket.OPEN;
  }

  /**
   * Send one typed message.
   * 发送一条类型化消息。
   */
  send(message: TSend): void {
    const data = this.serialize(message);
    this.socket.send(JSON.stringify(data));
    this.messagesSent += 1;
  }

  /**
   * Close the websocket connection.
   * 主动关闭 websocket 连接。
   */
  close(): void {
    this.status = 'closing';
    this.socket.close();
  }

  /**
   * Subscribe to all incoming messages.
   * 订阅所有接收到的消息。
   */
  onMessage(handler: (message: TReceive) => void): () => void {
    this.messageListeners.add(handler);
    return () => this.messageListeners.delete(handler);
  }

  /**
   * Subscribe to websocket open event.
   * 订阅 websocket 打开事件。
   */
  onOpen(handler: (event: Event) => void): () => void {
    this.openListeners.add(handler);
    return () => this.openListeners.delete(handler);
  }

  /**
   * Subscribe to websocket close event.
   * 订阅 websocket 关闭事件。
   */
  onClose(handler: (event: CloseEvent) => void): () => void {
    this.closeListeners.add(handler);
    return () => this.closeListeners.delete(handler);
  }

  /**
   * Subscribe to websocket error event.
   * 订阅 websocket 错误事件。
   */
  onError(handler: (event: Event) => void): () => void {
    this.errorListeners.add(handler);
    return () => this.errorListeners.delete(handler);
  }

  /**
   * Subscribe to messages by the `type` field.
   * 按消息的 `type` 字段进行订阅。
   */
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

  /**
   * Subscribe to typed payload messages with optional select/validate/decode steps.
   * 订阅类型化 payload 消息，并可通过 select/validate/decode 进行处理。
   */
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

// #endregion Typed WebSocket Client

// #region Endpoint Classes
// =====================================================

/**
 * WebSocket Product CRUD demo
 */
// Literal union is emitted as type because interface cannot model union values.
// 字面量联合类型使用 type，因为 interface 不能表达联合值。
export type ProductCrudWsDemoMessageType = string;
export interface ProductCrudWsDemoClientPayloadByType {
  create: ProductCreateRequest;
  delete: WsProductDeletePayload;
  list: ProductListQueryParams;
  update: WsProductUpdatePayload;
}
export type ProductCrudWsDemoSendUnion =
  | { type: 'create'; payload: ProductCreateRequest }
  | { type: 'delete'; payload: WsProductDeletePayload }
  | { type: 'list'; payload: ProductListQueryParams }
  | { type: 'update'; payload: WsProductUpdatePayload };
export class ProductCrudWsDemo<
  TSend = WebSocketMessage,
> extends TypedWebSocketClient<
  WebSocketMessage,
  TSend,
  ProductCrudWsDemoMessageType
> {
  static readonly NAME = 'productCrudWsDemo' as const;
  static readonly PATHS = {
    base: '/ws-go',
    group: '/v1',
    api: '/products-demo',
  } as const;
  static readonly FULL_PATH = '/ws-go/v1/products-demo' as const;
  static readonly MESSAGE_TYPES = [] as const;
  public readonly endpointName = ProductCrudWsDemo.NAME;
  public readonly endpointPath = ProductCrudWsDemo.FULL_PATH;

  constructor(options: WebSocketConvertOptions<TSend, WebSocketMessage>) {
    const url = ProductCrudWsDemo.FULL_PATH;
    super(url, options);
  }

  sendTypedMessage(message: ProductCrudWsDemoSendUnion): void {
    this.send(message as TSend);
  }
}
export function createProductCrudWsDemo<TSend = WebSocketMessage>(
  options: WebSocketConvertOptions<TSend, WebSocketMessage>
): ProductCrudWsDemo<TSend> {
  return new ProductCrudWsDemo<TSend>(options);
}

// #endregion Endpoint Classes
