// @flow
import * as React from 'react'
import Banner, {height} from './index'
import * as FsGen from '../../../actions/fs-gen'
import {namedConnect} from '../../../util/container'
import * as RowTypes from '../../row/types'

type OwnProps = {||}

const mapStateToProps = state => ({})

const mapDispatchToProps = dispatch => ({
  // TODO: set this to a "retry sync" action
  onRetry: () => dispatch(FsGen.createDriverEnable({})),
})

const ConnectedBanner = namedConnect<OwnProps, _, _, _, _>(
  mapStateToProps,
  mapDispatchToProps,
  (s, d, o) => ({...o, ...s, ...d}),
  'OutOfSpaceBanner'
)(Banner)

export default ConnectedBanner

export const asRows = (): Array<RowTypes.HeaderRowItem> => [
  {
    height: height,
    key: 'reset-banner',
    node: <ConnectedBanner />,
    rowType: 'header',
  },
]
