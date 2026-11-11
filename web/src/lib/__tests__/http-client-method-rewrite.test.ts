/*
Copyright (C) 2023-2026 QuantumNous

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as
published by the Free Software Foundation, either version 3 of the
License, or (at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program. If not, see <https://www.gnu.org/licenses/>.

For commercial licensing, please contact support@quantumnous.com
*/
import { describe, expect, it } from 'vitest'

import { rewriteBlockedWriteMethod } from '../http-client'

describe('rewriteBlockedWriteMethod', () => {
  it('maps PUT onto POST with an /update suffix', () => {
    expect(rewriteBlockedWriteMethod('/api/token/', 'put')).toEqual({
      url: '/api/token/update',
      method: 'post',
    })
    expect(rewriteBlockedWriteMethod('/api/channel/3', 'put')).toEqual({
      url: '/api/channel/3/update',
      method: 'post',
    })
  })

  it('maps PATCH onto POST with an /update suffix', () => {
    expect(rewriteBlockedWriteMethod('/api/option/model_pricing', 'patch')).toEqual(
      {
        url: '/api/option/model_pricing/update',
        method: 'post',
      }
    )
    expect(
      rewriteBlockedWriteMethod('/api/user/access_tokens/7', 'patch')
    ).toEqual({
      url: '/api/user/access_tokens/7/update',
      method: 'post',
    })
  })

  it('maps DELETE onto POST with a /delete suffix', () => {
    expect(rewriteBlockedWriteMethod('/api/token/5/', 'delete')).toEqual({
      url: '/api/token/5/delete',
      method: 'post',
    })
    expect(
      rewriteBlockedWriteMethod('/api/user/sessions/abc', 'delete')
    ).toEqual({
      url: '/api/user/sessions/abc/delete',
      method: 'post',
    })
  })

  it('keeps the query string after the suffix', () => {
    expect(
      rewriteBlockedWriteMethod('/api/token/?status_only=true', 'put')
    ).toEqual({
      url: '/api/token/update?status_only=true',
      method: 'post',
    })
    expect(
      rewriteBlockedWriteMethod('/api/performance/logs?mode=days', 'delete')
    ).toEqual({
      url: '/api/performance/logs/delete?mode=days',
      method: 'post',
    })
  })

  it('leaves GET/POST, non-console URLs and missing values untouched', () => {
    expect(rewriteBlockedWriteMethod('/api/token/5', 'get')).toEqual({
      url: '/api/token/5',
      method: 'get',
    })
    expect(rewriteBlockedWriteMethod('/api/token/5', 'post')).toEqual({
      url: '/api/token/5',
      method: 'post',
    })
    expect(rewriteBlockedWriteMethod('/v1/files/file-1', 'delete')).toEqual({
      url: '/v1/files/file-1',
      method: 'delete',
    })
    expect(rewriteBlockedWriteMethod(undefined, 'put')).toEqual({
      url: undefined,
      method: 'put',
    })
    expect(rewriteBlockedWriteMethod('/api/token/5', undefined)).toEqual({
      url: '/api/token/5',
      method: undefined,
    })
  })
})
