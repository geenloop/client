// @flow
import * as React from 'react'
import * as Kb from '../../../common-adapters'

type Props = {
  onRetry: () => void,
}

const Banner = (props: Props) => (
  <Kb.Banner
    text="You are out of storage space and some folders could not be properly synced. Make some space and"
    color="red"
    actions={[{onClick: props.onRetry, title: 'retry the sync.'}]}
  />
)

export default Banner
