import { useState } from 'react'

import { Button } from 'react-daisyui'

import ConditionTable from './ConditionTable'
import CreateForm from './CreateForm'
import EditForm from './EditForm'

import {
  GetProductDetailPageDataQuery,
  YahooAuctionCrawlSettingExcludeCondition,
} from '@/graphql/dist/client'

const YahooAuctionTabContent = ({ data }: { data: GetProductDetailPageDataQuery }) => {
  const [mode, setMode] = useState<'list' | 'create' | 'edit'>('list')
  const [condition, setCondition] = useState<YahooAuctionCrawlSettingExcludeCondition | undefined>(
    undefined,
  )

  return (
    <>
      <div className='space-y-4'>
        {mode == 'list' && (
          <>
            <Button onClick={() => setMode('create')} color='neutral' size='md' className='w-full'>
              ＋ 除外条件を追加
            </Button>
            <ConditionTable data={data} setMode={setMode} setCondition={setCondition} />
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
              ＜ 一覧に戻る
            </Button>
            <EditForm setMode={setMode} condition={condition} />
          </>
        )}
      </div>
    </>
  )
}

export default YahooAuctionTabContent
