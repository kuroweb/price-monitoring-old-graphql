'use client'

import { useState } from 'react'

import { Button } from 'react-daisyui'

import CreateForm from './CreateForm'
import EditForm from './EditForm'
import KeywordTable from './KeywordTable'

import {
  GetProductDetailPageDataQuery,
  IosysCrawlSettingExcludeKeyword,
} from '@/graphql/dist/client'

const IosysTab = ({ data }: { data: GetProductDetailPageDataQuery }) => {
  const [mode, setMode] = useState<'list' | 'create' | 'edit'>('list')
  const [condition, setCondition] = useState<IosysCrawlSettingExcludeKeyword | undefined>(undefined)

  return (
    <>
      <div className='space-y-4'>
        {mode == 'list' && (
          <>
            <Button onClick={() => setMode('create')} color='neutral' size='md' className='w-full'>
              ＋ 除外キーワードを追加
            </Button>
            <KeywordTable data={data} setMode={setMode} setCondition={setCondition} />
          </>
        )}
        {mode == 'create' && (
          <>
            <Button onClick={() => setMode('list')} color='neutral' size='md' className='w-full'>
              ＜ 一覧に戻る
            </Button>
            <CreateForm setMode={setMode} />
          </>
        )}
        {mode == 'edit' && (
          <>
            <Button onClick={() => setMode('list')} color='neutral' size='md' className='w-full'>
              ＜ 編集をキャンセルする
            </Button>
            <EditForm setMode={setMode} condition={condition} />
          </>
        )}
      </div>
    </>
  )
}

export default IosysTab
