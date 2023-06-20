import { useMutation } from '@apollo/client'
import { Button, FormControl, FormLabel, FormHelperText, FormErrorMessage, Stack, Input, Modal, ModalBody, ModalCloseButton, ModalContent, ModalHeader } from '@chakra-ui/react'
import { useForm, type SubmitHandler } from 'react-hook-form'

import { verifyVerificationCode as VERIFY_VERIFICATION_CODE } from '@gql'
import { usePropertyOnboarding } from '@hooks'

interface FormValues {
  verificationCode: string
}

interface Props {
  isOpen: boolean
  onClose: () => void
}

const VerificationModal = ({ isOpen, onClose }: Props): JSX.Element => {
  const { register, handleSubmit, formState: { errors } } = useForm<FormValues>()
  const [verifyCode, { loading: verifyingCode }] = useMutation(VERIFY_VERIFICATION_CODE)
  const { setStep, caretakerForm, setCaretakerVerified } = usePropertyOnboarding()
  
  // Verify phone
  const onSubmit: SubmitHandler<FormValues> = async data => {
    await verifyCode({
      variables: {
        input: {
          phone: `${caretakerForm.countryCode}${caretakerForm.phoneNumber}`,
          countryCode: "KE",
          verifyCode: data.verificationCode,
        },
      },
      // Proceed next step if successfull
      onCompleted: data => {
        const status = data?.verifyVerificationCode.success
        if (status === "approved") {
          setCaretakerVerified(true)
          setStep("units")
        }
      },
    })
  }

  return (
    <Modal isCentered isOpen={isOpen} onClose={onClose}>
      <ModalContent>
        <ModalHeader>Verify Phone</ModalHeader>
        <ModalCloseButton />
        <ModalBody>
          <form onSubmit={handleSubmit(onSubmit)}>
            <FormControl isInvalid={Boolean(errors?.verificationCode)}>
              <FormLabel>Enter Code</FormLabel>
              <Stack direction="row">
                <Input
                  {...register("verificationCode", { required: { value: true, message: "Invalid code" } })}
                  type="number"
                />
                <Button type="submit" isLoading={verifyingCode} disabled={verifyingCode} colorScheme="green">Verify</Button>
              </Stack>
              {(errors.verificationCode != null) && <FormErrorMessage>{errors?.verificationCode.message}</FormErrorMessage>}
              <FormHelperText>Enter 6-digit code sent to your phone</FormHelperText>
            </FormControl>
          </form>
        </ModalBody>
      </ModalContent>
    </Modal>
  )
}


export default VerificationModal
