// @flow
import {connect, type RouteProps} from '../../util/container'
import * as WalletsGen from '../../actions/wallets-gen'
import * as Constants from '../../constants/wallets'
import {anyWaiting} from '../../constants/waiting'
import HiddenString from '../../util/hidden-string'
import {Wrapper as LinkExisting} from '.'
import flags from '../../util/feature-flags'

type OwnProps = RouteProps<{backButton?: boolean, fromSendForm?: boolean, showOnCreation?: boolean}, {}>

const mapStateToProps = state => ({
  keyError: state.wallets.secretKeyError,
  linkExistingAccountError: state.wallets.linkExistingAccountError,
  nameError: state.wallets.accountNameError,
  nameValidationState: state.wallets.accountNameValidationState,
  secretKeyValidationState: state.wallets.secretKeyValidationState,
  waiting: anyWaiting(
    state,
    Constants.linkExistingWaitingKey,
    Constants.validateAccountNameWaitingKey,
    Constants.validateSecretKeyWaitingKey
  ),
})

const mapDispatchToProps = (dispatch, {navigateUp, routeProps}: OwnProps) => ({
  fromSendForm: routeProps.get('fromSendForm'),
  onCancel: () => navigateUp && dispatch(navigateUp()),
  onCheckKey: (key: string) => {
    dispatch(
      WalletsGen.createValidateSecretKey({
        secretKey: new HiddenString(key),
      })
    )
  },
  onCheckName: (name: string) => {
    dispatch(WalletsGen.createValidateAccountName({name}))
  },
  onClearErrors: () => dispatch(WalletsGen.createClearErrors()),
  onDone: (sk: string, name: string) => {
    dispatch(
      WalletsGen.createLinkExistingAccount({
        name,
        secretKey: new HiddenString(sk),
        setBuildingTo: routeProps.get('fromSendForm'),
        showOnCreation: routeProps.get('showOnCreation'),
      })
    )

    if (flags.useNewRouter) {
      dispatch(navigateUp())
    }
  },
})

const mergeProps = (stateProps, dispatchProps, ownProps) => ({
  keyError: stateProps.keyError,
  linkExistingAccountError: stateProps.linkExistingAccountError,
  nameError: stateProps.nameError,
  nameValidationState: stateProps.nameValidationState,
  onBack: ownProps.routeProps.get('backButton') ? dispatchProps.onCancel : undefined,
  onCancel: dispatchProps.onCancel,
  onCheckKey: dispatchProps.onCheckKey,
  onCheckName: dispatchProps.onCheckName,
  onClearErrors: dispatchProps.onClearErrors,
  onDone: dispatchProps.onDone,
  secretKeyValidationState: stateProps.secretKeyValidationState,
  waiting: stateProps.waiting,
})

export default connect<OwnProps, _, _, _, _>(
  mapStateToProps,
  mapDispatchToProps,
  mergeProps
)(LinkExisting)
