<template>
  <section class="crud-card">
    <div class="crud-header">
      <div>
        <h3>Product CRUD Demo (Http)</h3>
        <p class="crud-subtitle">
          Create, update, list, and delete products using GORM.
        </p>
      </div>
      <button class="btn secondary" type="button" @click="refresh">
        Refresh
      </button>
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
        <button class="btn ghost" type="submit">Create</button>
      </div>
    </form>

    <div v-if="error" class="error">{{ error }}</div>

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
          <button class="btn ghost" type="button" @click="update(item.id)">
            Update
          </button>
          <button class="btn danger" type="button" @click="remove(item.id)">
            Delete
          </button>
        </div>
      </div>
      <div v-if="items.length === 0" class="row empty">No products yet.</div>
    </div>
  </section>
</template>

<script setup lang="ts">
import axios from 'axios';

type EditState = {
  name: string;
  price: number;
  code: string;
  level: 'basic' | 'standard' | 'premium';
};

const items = ref<ProductModelResponse[]>([]);
const error = ref('');
const createForm = reactive<EditState>({
  name: '',
  price: 0,
  code: '',
  level: 'standard',
});
const edits = reactive<Record<number, EditState>>({});

const toUserMessage = (err: unknown, fallback: string): string => {
  if (axios.isAxiosError(err)) {
    if (!err.response) {
      return 'Cannot reach Gin API (http://127.0.0.1:8099). Start Go server or use `pnpm dev`.';
    }
    return `${fallback} (HTTP ${err.response.status})`;
  }
  return fallback;
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

const refresh = async () => {
  try {
    error.value = '';
    const res = await requestListProductsGet({
      query: { Page: 1, PageSize: 50 },
    });
    items.value = res.items;
    syncEdits(res.items);
  } catch (err) {
    console.error(err);
    error.value = toUserMessage(err, 'failed to load products');
  }
};

const create = async () => {
  try {
    error.value = '';
    if (!createForm.name || createForm.price <= 0) {
      error.value = 'name and price are required';
      return;
    }
    await requestCreateProductPost({
      name: createForm.name,
      price: createForm.price,
      code: createForm.code,
      level: createForm.level,
    });
    fillRandomCreateForm();
    await refresh();
  } catch (err) {
    console.error(err);
    error.value = toUserMessage(err, 'create failed');
  }
};

const update = async (id: number) => {
  try {
    error.value = '';
    const next = getEdit(id);
    const pathParams = { ID: String(id) };
    await requestUpdateProductPut(
      { path: pathParams },
      {
        name: next.name,
        price: next.price,
        code: next.code,
        level: next.level,
      }
    );
    await refresh();
  } catch (err) {
    console.error(err);
    error.value = toUserMessage(err, 'update failed');
  }
};

const remove = async (id: number) => {
  try {
    error.value = '';
    const pathParams = { ID: String(id) };
    await requestDeleteProductDelete({ path: pathParams });
    await refresh();
  } catch (err) {
    console.error(err);
    error.value = toUserMessage(err, 'delete failed');
  }
};

onMounted(() => {
  refresh();
  fillRandomCreateForm();
});
</script>

<style scoped lang="scss">
.crud-card {
  padding: 16px;
  margin: 10px 10px 10px 10px;
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

.field input {
  background: #0f1113;
  border: 1px solid rgba(255, 255, 255, 0.08);
  border-radius: 8px;
  padding: 8px 10px;
  color: #e5e7eb;
}

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

.btn.link {
  color: #60a5fa;
  font-weight: 600;
}

.btn.link:hover {
  text-decoration: underline;
}

.btn.secondary {
  background: #1f2937;
  color: #e5e7eb;
  border-color: rgba(255, 255, 255, 0.08);
}

.btn.ghost {
  background: transparent;
  color: #93c5fd;
}

.btn.danger {
  background: #b91c1c;
  color: #fff;
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
</style>
