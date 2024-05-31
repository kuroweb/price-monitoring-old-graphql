'use client'

import type { Dispatch, SetStateAction } from 'react'

import { useRouter, useParams, usePathname } from 'next/navigation'
import { Button } from 'react-daisyui'
import { useForm } from 'react-hook-form'
import { toast } from 'react-toastify'

import type { CreatePcKoubouCrawlSettingExcludeKeywordInput } from '@/graphql/dist/client'
import type { SubmitHandler } from 'react-hook-form'

import { createPcKoubouCrawlSettingExcludeKeyword } from '@/features/products/server-actions/pcKoubouCrawlSettingExcludeKeywordQuery'

const CreateForm = ({
  setMode,
}: {
  setMode: Dispatch<SetStateAction<'list' | 'create' | 'edit'>>
}) => {
  const router = useRouter()
  const params = useParams()
  const pathname = usePathname()

  const { register, handleSubmit } = useForm<CreatePcKoubouCrawlSettingExcludeKeywordInput>({
    defaultValues: {
      productId: String(params.id),
      keyword: null,
    },
  })

  const onSubmit: SubmitHandler<CreatePcKoubouCrawlSettingExcludeKeywordInput> = async (data) => {
    const result = await createPcKoubouCrawlSettingExcludeKeyword(data, pathname)
    if (result.data?.createPcKoubouCrawlSettingExcludeKeyword.ok) {
      toast.success('success')
      setMode('list')
    } else {
      toast.error('error')
    }
    router.refresh()
  }

  return (
    <>
      <form onSubmit={handleSubmit(onSubmit)} className='w-full space-y-2'>
        <label className='form-control'>
          <div className='label'>
            <span className='label-text'>キーワード</span>
          </div>
          <input {...register('keyword')} className='input input-bordered' />
        </label>
        <div className='pt-4'>
          <Button type='submit' color='primary' size='md' className='w-full'>
            追加
          </Button>
        </div>
      </form>
    </>
  )
}

export default CreateForm
