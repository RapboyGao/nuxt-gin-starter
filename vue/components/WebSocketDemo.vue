<template>
  <section class="crud-card">
    <div class="crud-header">
      <div>
        <h3>ProductCRUDDemo (WebSocket)</h3>
        <p class="crud-subtitle">
          Endpoint: <code>/ws-go/v1/products-demo</code>
        </p>
      </div>
      <div class="ws-toolbar">
        <div class="ws-status" :class="{ online: isOpen }">
          {{ isOpen ? 'Connected' : 'Disconnected' }}
        </div>
        <button
          class="btn secondary"
          type="button"
          @click="connect"
          :disabled="isOpen"
        >
          Connect
        </button>
        <button
          class="btn secondary"
          type="button"
          @click="disconnect"
          :disabled="!isOpen"
        >
          Disconnect
        </button>
        <button
          class="btn secondary"
          type="button"
          @click="requestList"
          :disabled="!isOpen"
        >
          Refresh
        </button>
      </div>
    </div>

    <form class="crud-form" @submit.prevent="create">
      <div class="field">
        <label>Name</label>
        <input v-model.trim="createForm.name" placeholder="e.g. Zen Chair" />
      </div>
      <div class="field">
        <label>Price</label>
        <input
          v-model.number="createForm.price"
          type="number"
          step="0.01"
          min="0"
        />
      </div>
      <div class="field">
        <label>Code</label>
        <input v-model.trim="createForm.code" placeholder="e.g. ZC-1001" />
      </div>
      <div class="field">
        <label>Level</label>
        <select v-model="createForm.level">
          <option value="basic">basic</option>
          <option value="standard">standard</option>
          <option value="premium">premium</option>
        </select>
      </div>
      <div class="field">
        <label></label>
        <button class="btn ghost" type="submit" :disabled="!isOpen">
          Create
        </button>
      </div>
    </form>

    <p v-if="error" class="error">{{ error }}</p>

    <div class="table">
      <div class="row head">
        <div>ID</div>
        <div>Name</div>
        <div>Price</div>
        <div>Code</div>
        <div>Level</div>
        <div>Actions</div>
      </div>
      <div v-for="item in items" :key="item.id" class="row">
        <div>{{ item.id }}</div>
        <div>
          <input v-model.trim="getEdit(item.id).name" />
        </div>
        <div>
          <input
            v-model.number="getEdit(item.id).price"
            type="number"
            step="0.01"
            min="0"
          />
        </div>
        <div>
          <input v-model.trim="getEdit(item.id).code" />
        </div>
        <div>
          <select v-model="getEdit(item.id).level">
            <option value="basic">basic</option>
            <option value="standard">standard</option>
            <option value="premium">premium</option>
          </select>
        </div>
        <div class="actions">
          <button
            class="btn ghost"
            type="button"
            @click="update(item.id)"
            :disabled="!isOpen"
          >
            Update
          </button>
          <button
            class="btn danger"
            type="button"
            @click="remove(item.id)"
            :disabled="!isOpen"
          >
            Delete
          </button>
        </div>
      </div>
      <div v-if="items.length === 0" class="row empty">No products yet.</div>
    </div>

    <div class="log">
      <div v-for="(item, i) in logs" :key="i" class="log-row">{{ item }}</div>
    </div>
  </section>
</template>

<script setup lang="ts">
import {
  ProductCrudWsDemo,
  createProductCrudWsDemo,
} from '@/composables/auto-generated-ws';
import { ensureWsProductServerEnvelope } from '@/composables/auto-generated-types';

type ProductWsClientMessage = {
  type: string;
  payload: unknown;
};

type EditState = {
  name: string;
  price: number;
  code: string;
  level: 'basic' | 'standard' | 'premium';
};

type ProductCrudWsClient = ProductCrudWsDemo<ProductWsClientMessage>;

const ws = shallowRef<ProductCrudWsClient | null>(null);
const isOpen = ref(false);
const error = ref('');
const logs = ref<string[]>([]);
const items = ref<ProductModelResponse[]>([]);
const createForm = reactive<EditState>({
  name: '',
  price: 0,
  code: '',
  level: 'standard',
});
const edits = reactive<Record<number, EditState>>({});
let unbindEvents: (() => void) | null = null;

const appendLog = (line: string) => {
  logs.value.unshift(line);
  if (logs.value.length > 40) {
    logs.value = logs.value.slice(0, 40);
  }
};

const randomName = () => {
  const names = [
    'Nova Lamp',
    'Echo Speaker',
    'Atlas Backpack',
    'Pixel Mug',
    'Zen Chair',
    'Orbit Watch',
    'Drift Keyboard',
    'Pulse Headphones',
    'Summit Bottle',
    'Glow Candle',
  ];
  return names[Math.floor(Math.random() * names.length)] ?? 'Product';
};

const randomPrice = () => Number((50 + Math.random() * 250).toFixed(2));

const randomCode = () => {
  const prefix = ['NX', 'ZG', 'PK', 'AT', 'GL'][Math.floor(Math.random() * 5)];
  const suffix = Math.floor(1000 + Math.random() * 9000);
  return `${prefix}-${suffix}`;
};

const randomLevel = (): 'basic' | 'standard' | 'premium' => {
  const levels: Array<'basic' | 'standard' | 'premium'> = [
    'basic',
    'standard',
    'premium',
  ];
  return levels[Math.floor(Math.random() * levels.length)] ?? 'standard';
};

const fillRandomCreateForm = () => {
  createForm.name = randomName();
  createForm.price = randomPrice();
  createForm.code = randomCode();
  createForm.level = randomLevel();
};

const syncEdits = (list: ProductModelResponse[]) => {
  for (const item of list) {
    edits[item.id] = {
      name: item.name,
      price: item.price,
      code: item.code,
      level: item.level,
    };
  }
};

const getEdit = (id: number) => {
  if (!edits[id]) {
    edits[id] = {
      name: '',
      price: 0,
      code: '',
      level: 'standard',
    };
  }
  return edits[id];
};

const applyList = (list: ProductModelResponse[]) => {
  items.value = list;
  syncEdits(list);
};

const createClient = () =>
  createProductCrudWsDemo<ProductWsClientMessage>({
    serialize: (value) => value,
    deserialize: (value) => ensureWsProductServerEnvelope(value),
  });

const requestList = () => {
  if (!ws.value || !isOpen.value) return;
  ws.value.sendListPayload({ page: 1, pageSize: 0 });
};

const create = () => {
  if (!isOpen.value) return;
  if (!createForm.name || createForm.price <= 0) {
    error.value = 'name and price are required';
    return;
  }
  error.value = '';
  ws.value?.send({
    type: 'create',
    payload: {
    name: createForm.name,
    price: createForm.price,
    code: createForm.code,
    level: createForm.level,
    },
  });
  fillRandomCreateForm();
};

const update = (id: number) => {
  if (!isOpen.value) return;
  const next = getEdit(id);
  error.value = '';
  ws.value?.send({
    type: 'update',
    payload: {
      id,
      name: next.name,
      price: next.price,
      code: next.code,
      level: next.level,
    },
  });
};

const remove = (id: number) => {
  if (!isOpen.value) return;
  error.value = '';
  ws.value?.send({
    type: 'delete',
    payload: { id },
  });
};

const bindClientEvents = (client: ProductCrudWsClient) => {
  unbindEvents?.();

  const offOpen = client.onOpen(() => {
    isOpen.value = true;
    appendLog('connected');
    requestList();
  });

  const offClose = client.onClose(() => {
    isOpen.value = false;
    appendLog('disconnected');
  });

  const offError = client.onError(() => {
    error.value = 'websocket error';
  });

  const offList = client.onListType((message) => {
    applyList(message.items ?? []);
  });

  const offSync = client.onSyncType((message) => {
    applyList(message.items ?? []);
  });

  const offCreated = client.onCreatedType((message) => {
    appendLog(`[created] ${message.item?.name ?? ''}`);
  });

  const offUpdated = client.onUpdatedType((message) => {
    appendLog(`[updated] ${message.item?.name ?? ''}`);
  });

  const offDeleted = client.onDeletedType((message) => {
    appendLog(`[deleted] ${message.deletedId}`);
  });

  const offSystem = client.onSystemType((message) => {
    appendLog(`[system] ${message.message ?? ''}`);
  });

  const offServerError = client.onErrorType((message) => {
    error.value = message.message || 'server error';
    appendLog(`[error] ${message.message ?? ''}`);
  });

  unbindEvents = () => {
    offOpen();
    offClose();
    offError();
    offList();
    offSync();
    offCreated();
    offUpdated();
    offDeleted();
    offSystem();
    offServerError();
  };
};

const connect = () => {
  error.value = '';
  if (ws.value?.isOpen || ws.value?.readyState === WebSocket.CONNECTING) {
    return;
  }
  ws.value = createClient();
  bindClientEvents(ws.value);
};

const disconnect = () => {
  ws.value?.close();
};

onMounted(() => {
  fillRandomCreateForm();
  connect();
});

onBeforeUnmount(() => {
  disconnect();
  unbindEvents?.();
  unbindEvents = null;
});
</script>

<style scoped lang="scss">
.crud-card {
  padding: 16px;
  margin: 10px;
  border-radius: 14px;
  background: #111214;
  border: 1px solid rgba(255, 255, 255, 0.08);
  box-shadow: 0 10px 24px rgba(0, 0, 0, 0.2);
  color: #e8e8e8;
}

.crud-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 12px;
}

.crud-header h3 {
  margin: 0;
  font-size: 16px;
}

.crud-subtitle {
  margin: 4px 0 0;
  color: #9aa3af;
  font-size: 12px;
}

.ws-toolbar {
  display: flex;
  align-items: center;
  gap: 8px;
  flex-wrap: wrap;
  justify-content: flex-end;
}

.ws-status {
  font-size: 12px;
  color: #ef4444;
  min-width: 92px;
  text-align: right;
}

.ws-status.online {
  color: #22c55e;
}

.crud-form {
  margin-top: 14px;
  display: grid;
  gap: 12px;
  grid-template-columns: repeat(auto-fit, minmax(160px, 1fr));
}

.field {
  display: grid;
  gap: 6px;
}

.field label {
  font-size: 12px;
  color: #9aa3af;
}

.field input,
.field select {
  background: #0f1113;
  border: 1px solid rgba(255, 255, 255, 0.08);
  border-radius: 8px;
  padding: 8px 10px;
  color: #e5e7eb;
}

.btn {
  border: none;
  background: transparent;
  padding: 0;
  font-size: 12px;
  cursor: pointer;
}

.btn.secondary {
  background: #1f2937;
  color: #e5e7eb;
  border-color: rgba(255, 255, 255, 0.08);
  border-radius: 8px;
  padding: 7px 12px;
}

.btn.ghost {
  background: transparent;
  color: #93c5fd;
}

.btn.danger {
  background: #b91c1c;
  color: #fff;
  border-radius: 4px;
  padding: 0 2px;
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

.table {
  margin-top: 14px;
  display: grid;
  gap: 8px;
}

.row {
  display: grid;
  gap: 8px;
  grid-template-columns: 60px minmax(220px, 1fr) 100px 120px 100px 150px;
  align-items: center;
  padding: 8px;
  border-radius: 10px;
  background: rgba(255, 255, 255, 0.03);
  border: 1px solid rgba(255, 255, 255, 0.04);
  font-size: 12px;
}

.row.head {
  background: rgba(255, 255, 255, 0.06);
  font-weight: 600;
  color: #d1d5db;
}

.row.empty {
  text-align: center;
  color: #6b7280;
}

.row input,
.row select {
  width: 100%;
  background: #0f1113;
  border: 1px solid rgba(255, 255, 255, 0.08);
  border-radius: 6px;
  padding: 6px 8px;
  color: #e5e7eb;
}

.actions {
  display: flex;
  gap: 8px;
  flex-wrap: nowrap;
  white-space: nowrap;
  align-items: center;
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
