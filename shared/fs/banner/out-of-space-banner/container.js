// @flow
import Banner from './index'
import * as FsGen from '../../../actions/fs-gen'
import {namedConnect} from '../../../util/container'

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
