/**
 * =====================================================
 * Nuxt Gin Shared Schemas
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

// #region Shared Helpers
// =====================================================

const isPlainObject = (value: unknown): value is Record<string, unknown> =>
  Object.prototype.toString.call(value) === '[object Object]';

// #endregion Shared Helpers

// #region Interfaces & Validators
// =====================================================

// #region Interfaces & Validators
// =====================================================

// =====================================================
// INTERFACES & VALIDATORS
// Default: object schemas use interface.
// Fallback: use type only when interface cannot model the shape.
// 默认：对象结构使用 interface。
// 兜底：只有 interface 无法表达时才使用 type。
// =====================================================

// -----------------------------------------------------
// TYPE: ProductCreateRequest
// -----------------------------------------------------
export interface ProductCreateRequest {
  /** Product name */
  name: string;
  /** Product unit price, must be greater than 0 */
  price: number;
  /** Product unique code */
  code: string;
  /** Product level */
  level: 'basic' | 'standard' | 'premium';
}

/**
 * Validate whether a value matches ProductCreateRequest.
 * 校验一个值是否符合 ProductCreateRequest 结构。
 */

/**
 * Validate whether a value matches ProductCreateRequest.
 * 校验一个值是否符合 ProductCreateRequest 结构。
 */
export function validateProductCreateRequest(
  value: unknown
): value is ProductCreateRequest {
  if (!isPlainObject(value)) return false;
  const obj = value as Record<string, unknown>;
  if (!('name' in obj)) return false;
  if (!(typeof obj['name'] === 'string')) return false;
  if (!('price' in obj)) return false;
  if (!(typeof obj['price'] === 'number')) return false;
  if (!('code' in obj)) return false;
  if (!(typeof obj['code'] === 'string')) return false;
  if (!('level' in obj)) return false;
  if (
    !(
      typeof obj['level'] === 'string' &&
      (obj['level'] === 'basic' ||
        obj['level'] === 'standard' ||
        obj['level'] === 'premium')
    )
  )
    return false;
  return true;
}

// -----------------------------------------------------
// TYPE: ProductModelResponse
// -----------------------------------------------------

// -----------------------------------------------------
// TYPE: ProductModelResponse
// -----------------------------------------------------
export interface ProductModelResponse {
  /** Product primary key */
  id: number;
  /** Product name */
  name: string;
  /** Product unit price */
  price: number;
  /** Product unique code */
  code: string;
  /** Product level */
  level: 'basic' | 'standard' | 'premium';
  /** Creation timestamp in milliseconds */
  createdAt: number;
  /** Last update timestamp in milliseconds */
  updatedAt: number;
}

/**
 * Validate whether a value matches ProductModelResponse.
 * 校验一个值是否符合 ProductModelResponse 结构。
 */

/**
 * Validate whether a value matches ProductModelResponse.
 * 校验一个值是否符合 ProductModelResponse 结构。
 */
export function validateProductModelResponse(
  value: unknown
): value is ProductModelResponse {
  if (!isPlainObject(value)) return false;
  const obj = value as Record<string, unknown>;
  if (!('id' in obj)) return false;
  if (!(typeof obj['id'] === 'number')) return false;
  if (!('name' in obj)) return false;
  if (!(typeof obj['name'] === 'string')) return false;
  if (!('price' in obj)) return false;
  if (!(typeof obj['price'] === 'number')) return false;
  if (!('code' in obj)) return false;
  if (!(typeof obj['code'] === 'string')) return false;
  if (!('level' in obj)) return false;
  if (
    !(
      typeof obj['level'] === 'string' &&
      (obj['level'] === 'basic' ||
        obj['level'] === 'standard' ||
        obj['level'] === 'premium')
    )
  )
    return false;
  if (!('createdAt' in obj)) return false;
  if (!(typeof obj['createdAt'] === 'number')) return false;
  if (!('updatedAt' in obj)) return false;
  if (!(typeof obj['updatedAt'] === 'number')) return false;
  return true;
}

// -----------------------------------------------------
// TYPE: ProductIDPathParams
// -----------------------------------------------------

// -----------------------------------------------------
// TYPE: ProductIDPathParams
// -----------------------------------------------------
export interface ProductIDPathParams {
  /** Product identifier in route path */
  ID: string;
}

/**
 * Validate whether a value matches ProductIDPathParams.
 * 校验一个值是否符合 ProductIDPathParams 结构。
 */

/**
 * Validate whether a value matches ProductIDPathParams.
 * 校验一个值是否符合 ProductIDPathParams 结构。
 */
export function validateProductIDPathParams(
  value: unknown
): value is ProductIDPathParams {
  if (!isPlainObject(value)) return false;
  const obj = value as Record<string, unknown>;
  if (!('ID' in obj)) return false;
  if (!(typeof obj['ID'] === 'string')) return false;
  return true;
}

// -----------------------------------------------------
// TYPE: ProductUpdateRequest
// -----------------------------------------------------

// -----------------------------------------------------
// TYPE: ProductUpdateRequest
// -----------------------------------------------------
export interface ProductUpdateRequest {
  /** Product name, optional in partial update */
  name: string;
  /** Product unit price, optional in partial update */
  price: number;
  /** Product code, optional in partial update */
  code: string;
  /** Product level, optional in partial update */
  level: 'basic' | 'standard' | 'premium';
}

/**
 * Validate whether a value matches ProductUpdateRequest.
 * 校验一个值是否符合 ProductUpdateRequest 结构。
 */

/**
 * Validate whether a value matches ProductUpdateRequest.
 * 校验一个值是否符合 ProductUpdateRequest 结构。
 */
export function validateProductUpdateRequest(
  value: unknown
): value is ProductUpdateRequest {
  if (!isPlainObject(value)) return false;
  const obj = value as Record<string, unknown>;
  if (!('name' in obj)) return false;
  if (!(typeof obj['name'] === 'string')) return false;
  if (!('price' in obj)) return false;
  if (!(typeof obj['price'] === 'number')) return false;
  if (!('code' in obj)) return false;
  if (!(typeof obj['code'] === 'string')) return false;
  if (!('level' in obj)) return false;
  if (
    !(
      typeof obj['level'] === 'string' &&
      (obj['level'] === 'basic' ||
        obj['level'] === 'standard' ||
        obj['level'] === 'premium')
    )
  )
    return false;
  return true;
}

// -----------------------------------------------------
// TYPE: ProductListQueryParams
// -----------------------------------------------------

// -----------------------------------------------------
// TYPE: ProductListQueryParams
// -----------------------------------------------------
export interface ProductListQueryParams {
  /** Page number, starting from 1 */
  Page: number;
  /** Page size, max 100 */
  PageSize: number;
}

/**
 * Validate whether a value matches ProductListQueryParams.
 * 校验一个值是否符合 ProductListQueryParams 结构。
 */

/**
 * Validate whether a value matches ProductListQueryParams.
 * 校验一个值是否符合 ProductListQueryParams 结构。
 */
export function validateProductListQueryParams(
  value: unknown
): value is ProductListQueryParams {
  if (!isPlainObject(value)) return false;
  const obj = value as Record<string, unknown>;
  if (!('Page' in obj)) return false;
  if (!(typeof obj['Page'] === 'number')) return false;
  if (!('PageSize' in obj)) return false;
  if (!(typeof obj['PageSize'] === 'number')) return false;
  return true;
}

// -----------------------------------------------------
// TYPE: ProductListResponse
// -----------------------------------------------------

// -----------------------------------------------------
// TYPE: ProductListResponse
// -----------------------------------------------------
export interface ProductListResponse {
  /** Current page product items */
  items: ProductModelResponse[];
  /** Total item count */
  total: number;
  /** Current page number */
  page: number;
  /** Current page size */
  size: number;
}

/**
 * Validate whether a value matches ProductListResponse.
 * 校验一个值是否符合 ProductListResponse 结构。
 */

/**
 * Validate whether a value matches ProductListResponse.
 * 校验一个值是否符合 ProductListResponse 结构。
 */
export function validateProductListResponse(
  value: unknown
): value is ProductListResponse {
  if (!isPlainObject(value)) return false;
  const obj = value as Record<string, unknown>;
  if (!('items' in obj)) return false;
  if (
    !(
      Array.isArray(obj['items']) &&
      obj['items'].every((v1) => validateProductModelResponse(v1))
    )
  )
    return false;
  if (!('total' in obj)) return false;
  if (!(typeof obj['total'] === 'number')) return false;
  if (!('page' in obj)) return false;
  if (!(typeof obj['page'] === 'number')) return false;
  if (!('size' in obj)) return false;
  if (!(typeof obj['size'] === 'number')) return false;
  return true;
}

// #endregion Interfaces & Validators

// #region Interfaces & Validators
// =====================================================

// =====================================================
// INTERFACES & VALIDATORS
// Default: object schemas use interface.
// Fallback: use type only when interface cannot model the shape.
// 默认：对象结构使用 interface。
// 兜底：只有 interface 无法表达时才使用 type。
// =====================================================

// -----------------------------------------------------
// TYPE: WsProductClientMessage
// -----------------------------------------------------
export interface WsProductClientMessage {
  /** WebSocket action type */
  type: string;
  /** WebSocket action payload */
  payload: Record<string, unknown>;
}

/**
 * Validate whether a value matches WsProductClientMessage.
 * 校验一个值是否符合 WsProductClientMessage 结构。
 */

/**
 * Validate whether a value matches WsProductClientMessage.
 * 校验一个值是否符合 WsProductClientMessage 结构。
 */
export function validateWsProductClientMessage(
  value: unknown
): value is WsProductClientMessage {
  if (!isPlainObject(value)) return false;
  const obj = value as Record<string, unknown>;
  if (!('type' in obj)) return false;
  if (!(typeof obj['type'] === 'string')) return false;
  if (!('payload' in obj)) return false;
  if (
    !(
      isPlainObject(obj['payload']) &&
      Object.values(obj['payload']).every((v1) => true)
    )
  )
    return false;
  return true;
}

/**
 * Ensure a typed WsProductClientMessage after validation.
 * 先校验，再确保得到类型化的 WsProductClientMessage。
 */

/**
 * Ensure a typed WsProductClientMessage after validation.
 * 先校验，再确保得到类型化的 WsProductClientMessage。
 */
export function ensureWsProductClientMessage(
  value: unknown
): WsProductClientMessage {
  if (!validateWsProductClientMessage(value)) {
    throw new Error('Invalid WsProductClientMessage');
  }
  return value;
}

// -----------------------------------------------------
// TYPE: ProductModelResponse
// -----------------------------------------------------

/**
 * Ensure a typed ProductModelResponse after validation.
 * 先校验，再确保得到类型化的 ProductModelResponse。
 */
export function ensureProductModelResponse(
  value: unknown
): ProductModelResponse {
  if (!validateProductModelResponse(value)) {
    throw new Error('Invalid ProductModelResponse');
  }
  return value;
}

// -----------------------------------------------------
// TYPE: WsProductServerEnvelope
// -----------------------------------------------------

// -----------------------------------------------------
// TYPE: WsProductServerEnvelope
// -----------------------------------------------------
export interface WsProductServerEnvelope {
  /** Server event type */
  type: string;
  /** Event message */
  message: string;
  /** Single product payload */
  item: ProductModelResponse;
  /** Product list payload */
  items: ProductModelResponse[];
  /** Total item count */
  total: number;
  /** Current page number */
  page: number;
  /** Current page size */
  size: number;
  /** Deleted product id */
  deletedId: number;
  /** Server timestamp in milliseconds */
  at: number;
}

/**
 * Validate whether a value matches WsProductServerEnvelope.
 * 校验一个值是否符合 WsProductServerEnvelope 结构。
 */

/**
 * Validate whether a value matches WsProductServerEnvelope.
 * 校验一个值是否符合 WsProductServerEnvelope 结构。
 */
export function validateWsProductServerEnvelope(
  value: unknown
): value is WsProductServerEnvelope {
  if (!isPlainObject(value)) return false;
  const obj = value as Record<string, unknown>;
  if (!('type' in obj)) return false;
  if (!(typeof obj['type'] === 'string')) return false;
  if (!('message' in obj)) return false;
  if (!(typeof obj['message'] === 'string')) return false;
  if (!('item' in obj)) return false;
  if (!validateProductModelResponse(obj['item'])) return false;
  if (!('items' in obj)) return false;
  if (
    !(
      Array.isArray(obj['items']) &&
      obj['items'].every((v1) => validateProductModelResponse(v1))
    )
  )
    return false;
  if (!('total' in obj)) return false;
  if (!(typeof obj['total'] === 'number')) return false;
  if (!('page' in obj)) return false;
  if (!(typeof obj['page'] === 'number')) return false;
  if (!('size' in obj)) return false;
  if (!(typeof obj['size'] === 'number')) return false;
  if (!('deletedId' in obj)) return false;
  if (!(typeof obj['deletedId'] === 'number')) return false;
  if (!('at' in obj)) return false;
  if (!(typeof obj['at'] === 'number')) return false;
  return true;
}

/**
 * Ensure a typed WsProductServerEnvelope after validation.
 * 先校验，再确保得到类型化的 WsProductServerEnvelope。
 */

/**
 * Ensure a typed WsProductServerEnvelope after validation.
 * 先校验，再确保得到类型化的 WsProductServerEnvelope。
 */
export function ensureWsProductServerEnvelope(
  value: unknown
): WsProductServerEnvelope {
  if (!validateWsProductServerEnvelope(value)) {
    throw new Error('Invalid WsProductServerEnvelope');
  }
  return value;
}

// #endregion Interfaces & Validators

// #endregion Interfaces & Validators
