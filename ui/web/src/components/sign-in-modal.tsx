import { Modal, ModalBody, ModalContent, ModalCloseButton, ModalHeader } from '@chakra-ui/react'
import { useSignIn } from '@hooks'

import SignInForm from 'form/sign-in'
import VerifySignInForm from 'form/verify-signin'

interface Props {
	isOpen: boolean
	onClose: () => void
}

const SignInModal = ({ isOpen, onClose }: Props): JSX.Element => {
	const { status } = useSignIn()

	return (
		<Modal isCentered isOpen={isOpen} onClose={onClose}>
			<ModalContent>
				<ModalHeader>{status !== "approved" ? `Sign In with Phone` : `Verify Phone`}</ModalHeader>
				<ModalCloseButton />
				<ModalBody p={5}>
					{status === 'sign-in' && <SignInForm />}
					{status === 'pending' && <VerifySignInForm />}
				</ModalBody>
			</ModalContent>
		</Modal>
	)
}

export default SignInModal
