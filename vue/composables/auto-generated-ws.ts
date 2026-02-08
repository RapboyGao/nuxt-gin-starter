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

const isoDateLike = /^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(?:\.\d{1,9})?(?:Z|[+\-]\d{2}:\d{2})$/;

const normalizeRequestJSON = (value: unknown): unknown => {
  if (value instanceof Date) return value.toISOString();
  if (Array.isArray(value)) return value.map(normalizeRequestJSON);
  if (isPlainObject(value)) {
    const out: Record<string, unknown> = {};
    for (const [k, v] of Object.entries(value)) out[k] = normalizeRequestJSON(v);
    return out;
  }
  return value;
};

const normalizeResponseJSON = (value: unknown): unknown => {
  if (Array.isArray(value)) return value.map(normalizeResponseJSON);
  if (typeof value === 'string' && isoDateLike.test(value)) {
    const date = new Date(value);
    if (!Number.isNaN(date.getTime())) return date;
  }
  if (isPlainObject(value)) {
    const out: Record<string, unknown> = {};
    for (const [k, v] of Object.entries(value)) out[k] = normalizeResponseJSON(v);
    return out;
  }
  return value;
};

export interface WebSocketConvertOptions<TSend = unknown, TReceive = unknown> {
  serialize?: (value: TSend) => unknown;
  deserialize?: (value: unknown) => TReceive;
}

export interface WebSocketClient<TReceive = unknown, TSend = unknown> {
  socket: WebSocket;
  send: (message: TSend) => void;
  close: () => void;
  onMessage: (handler: (message: TReceive) => void) => () => void;
  onOpen: (handler: (event: Event) => void) => () => void;
  onClose: (handler: (event: CloseEvent) => void) => () => void;
  onError: (handler: (event: Event) => void) => () => void;
}

const resolveWebSocketURL = (url: string): string => {
  if (url.startsWith('ws://') || url.startsWith('wss://')) return url;
  if (url.startsWith('http://')) return `ws://${url.slice(7)}`;
  if (url.startsWith('https://')) return `wss://${url.slice(8)}`;
  if (url.startsWith('/')) {
    const isHttps = typeof window !== 'undefined' && window.location?.protocol === 'https:';
    const host = typeof window !== 'undefined' ? window.location.host : '';
    return `${isHttps ? 'wss' : 'ws'}://${host}${url}`;
  }
  return url;
};

const joinURLPath = (baseURL: string, path: string): string => {
  const base = baseURL.trim();
  const p = path.trim();
  if (!base) return p.startsWith('/') ? p : `/${p}`;
  if (!p) return base.startsWith('/') ? base.replace(/\/+$/, '') : `/${base.replace(/\/+$/, '')}`;
  const trimmedBase = base.replace(/\/+$/, '');
  const trimmedPath = p.replace(/^\/+/, '');
  return trimmedBase.startsWith('/') ? `${trimmedBase}/${trimmedPath}` : `/${trimmedBase}/${trimmedPath}`;
};

const createWebSocketClient = <TReceive, TSend>(
  url: string,
  options?: WebSocketConvertOptions<TSend, TReceive>
): WebSocketClient<TReceive, TSend> => {
  const socket = new WebSocket(resolveWebSocketURL(url));
  const messageListeners = new Set<(message: TReceive) => void>();
  const openListeners = new Set<(event: Event) => void>();
  const closeListeners = new Set<(event: CloseEvent) => void>();
  const errorListeners = new Set<(event: Event) => void>();
  const serialize = options?.serialize ?? ((value: TSend) => normalizeRequestJSON(value));
  const deserialize = options?.deserialize ?? ((value: unknown) => normalizeResponseJSON(value) as TReceive);

  socket.addEventListener('message', (event) => {
    let payload: unknown = event.data;
    if (typeof payload === 'string') {
      try {
        payload = JSON.parse(payload);
      } catch {
        // keep raw payload
      }
    }
    const message = deserialize(payload);
    for (const listener of messageListeners) listener(message);
  });
  socket.addEventListener('open', (event) => {
    for (const listener of openListeners) listener(event);
  });
  socket.addEventListener('close', (event) => {
    for (const listener of closeListeners) listener(event);
  });
  socket.addEventListener('error', (event) => {
    for (const listener of errorListeners) listener(event);
  });

  return {
    socket,
    send: (message: TSend) => {
      const data = serialize(message);
      socket.send(JSON.stringify(data));
    },
    close: () => socket.close(),
    onMessage: (handler) => {
      messageListeners.add(handler);
      return () => messageListeners.delete(handler);
    },
    onOpen: (handler) => {
      openListeners.add(handler);
      return () => openListeners.delete(handler);
    },
    onClose: (handler) => {
      closeListeners.add(handler);
      return () => closeListeners.delete(handler);
    },
    onError: (handler) => {
      errorListeners.add(handler);
      return () => errorListeners.delete(handler);
    },
  };
};

export interface WebSocketMessage {
  type: string;
  payload: string;
}

export interface WsServerEnvelope {
  type: string;
  client: string;
  message: string;
  at: string;
}

/**
 * WebSocket demo with typed message handlers
 */
export function chatDemo(options?: WebSocketConvertOptions<WebSocketMessage, WsServerEnvelope>): WebSocketClient<WsServerEnvelope, WebSocketMessage> {
  const url = joinURLPath('/ws-go/v1', '/chat-demo');
  return createWebSocketClient<WsServerEnvelope, WebSocketMessage>(url, options);
}
