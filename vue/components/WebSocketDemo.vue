<template>
  <section class="ws-card">
    <div class="ws-header">
      <div>
        <h3>WebSocket Demo</h3>
        <p class="ws-subtitle">Endpoint: <code>/ws-go/v1/chat-demo</code></p>
      </div>
      <div class="ws-status" :class="{ online: isOpen }">
        {{ isOpen ? 'Connected' : 'Disconnected' }}
      </div>
    </div>

    <div class="ws-actions">
      <button class="btn" type="button" @click="connect" :disabled="isOpen">
        Connect
      </button>
      <button class="btn" type="button" @click="disconnect" :disabled="!isOpen">
        Disconnect
      </button>
      <button class="btn" type="button" @click="sendWhoAmI" :disabled="!isOpen">
        WhoAmI
      </button>
      <button class="btn" type="button" @click="sendPing" :disabled="!isOpen">
        Ping
      </button>
    </div>

    <form class="chat-form" @submit.prevent="sendChat">
      <input v-model.trim="user" placeholder="user" />
      <input v-model.trim="chatText" placeholder="message" />
      <button class="btn primary" type="submit" :disabled="!isOpen">
        Send Chat
      </button>
    </form>

    <p v-if="error" class="error">{{ error }}</p>

    <div class="log">
      <div v-for="(item, i) in logs" :key="i" class="log-row">{{ item }}</div>
    </div>
  </section>
</template>

<script setup lang="ts">
type WsClientMessage = {
  type: string;
  payload: unknown;
};

type ChatDemoClient = ChatDemo<WsClientMessage>;

const ws = shallowRef<ChatDemoClient | null>(null);
const isOpen = ref(false);
const error = ref('');
const randomUser = () => {
  const pool = ['nova', 'echo', 'atlas', 'pixel', 'zen', 'orbit', 'pulse'];
  const suffix = Math.floor(1000 + Math.random() * 9000);
  return `${pool[Math.floor(Math.random() * pool.length)]}-${suffix}`;
};
const user = ref(randomUser());
const chatText = ref('');
const logs = ref<string[]>([]);
let unbindEvents: (() => void) | null = null;

const createClient = () => {
  let demo = createChatDemo<WsClientMessage>({
    serialize: (value) => value,
    deserialize: (value) => ensureWsServerEnvelope(value),
  });
  console.log('createClient', demo);
  return demo;
};
const appendLog = (line: string) => {
  logs.value.unshift(line);
  if (logs.value.length > 30) {
    logs.value = logs.value.slice(0, 30);
  }
};

const bindClientEvents = (client: ChatDemoClient) => {
  unbindEvents?.();

  const offOpen = client.onOpen(() => {
    isOpen.value = true;
    appendLog('connected');
  });

  const offMessage = client.onMessage((payload) => {
    appendLog(`[${payload.type}] ${payload.message ?? ''}`);
  });

  const offChat = client.onChatType((payload) => {
    appendLog(`[chat] ${payload.message ?? ''}`);
  });
  const offWhoami = client.onWhoamiType((payload) => {
    appendLog(`[whoami] ${payload.client ?? ''}`);
  });
  const offPong = client.onPongType((payload) => {
    appendLog(`[pong] ${payload.message ?? ''}`);
  });
  const offSystem = client.onSystemType((payload) => {
    appendLog(`[system] ${payload.message ?? ''}`);
  });
  const offServerError = client.onErrorType((payload) => {
    appendLog(`[error] ${payload.message ?? ''}`);
  });

  const offError = client.onError(() => {
    error.value = 'websocket error';
  });

  const offClose = client.onClose(() => {
    isOpen.value = false;
    appendLog('disconnected');
  });

  unbindEvents = () => {
    offOpen();
    offMessage();
    offChat();
    offWhoami();
    offPong();
    offSystem();
    offServerError();
    offError();
    offClose();
  };
};

const connect = () => {
  error.value = '';
  if (ws.value?.isOpen || ws.value?.readyState === WebSocket.CONNECTING) {
    return;
  }
  if (!ws.value || ws.value.readyState >= WebSocket.CLOSING) {
    ws.value = createClient();
    bindClientEvents(ws.value);
  }
};

const disconnect = () => {
  ws.value?.close();
};

const send = (type: string, payload: unknown) => {
  if (!ws.value || !isOpen.value) return;
  ws.value.send({ type, payload });
};

const sendWhoAmI = () => {
  send('whoami', {});
};

const sendPing = () => {
  send('ping', { at: Date.now() });
};

const sendChat = () => {
  const text = chatText.value.trim();
  if (!text) return;
  send('chat', { user: user.value.trim() || 'demo-user', content: text });
  chatText.value = '';
};

onMounted(() => {
  ws.value = createClient();
  bindClientEvents(ws.value);
  connect();
});

onBeforeUnmount(() => {
  disconnect();
  unbindEvents?.();
  unbindEvents = null;
});
</script>

<style scoped lang="scss">
.ws-card {
  margin: 10px 10px 10px 10px;
  padding: 16px;
  border-radius: 14px;
  background: #111214;
  border: 1px solid rgba(255, 255, 255, 0.08);
  color: #e8e8e8;
}

.ws-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 12px;
}

.ws-header h3 {
  margin: 0;
}

.ws-subtitle {
  margin: 4px 0 0;
  color: #9ca3af;
  font-size: 12px;
}

.ws-status {
  font-size: 12px;
  color: #ef4444;
}

.ws-status.online {
  color: #22c55e;
}

.ws-actions {
  margin-top: 12px;
  display: flex;
  gap: 8px;
  flex-wrap: wrap;
}

.chat-form {
  margin-top: 12px;
  display: grid;
  gap: 8px;
  grid-template-columns: 140px minmax(140px, 1fr) 110px;
}

.chat-form input {
  background: #0b0c0e;
  border: 1px solid rgba(255, 255, 255, 0.1);
  border-radius: 8px;
  padding: 7px 9px;
  color: #e8e8e8;
}

.btn {
  border: 1px solid rgba(255, 255, 255, 0.12);
  background: #1f2937;
  color: #e5e7eb;
  border-radius: 8px;
  padding: 6px 10px;
  cursor: pointer;
  font-size: 12px;
}

.btn.primary {
  background: #2563eb;
  border-color: #2563eb;
}

.btn:disabled {
  opacity: 0.6;
  cursor: not-allowed;
}

.error {
  margin-top: 10px;
  color: #fca5a5;
  font-size: 12px;
}

.log {
  margin-top: 12px;
  border-radius: 10px;
  border: 1px solid rgba(255, 255, 255, 0.08);
  background: rgba(255, 255, 255, 0.02);
  max-height: 220px;
  overflow: auto;
}

.log-row {
  padding: 7px 10px;
  border-bottom: 1px solid rgba(255, 255, 255, 0.05);
  font-size: 12px;
  color: #cbd5e1;
}

.log-row:last-child {
  border-bottom: none;
}
</style>
