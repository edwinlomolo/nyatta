'use client'

import { useMutation } from '@apollo/client'
import { Center, FormControl, FormLabel, Input, Icon, Image, Button, Spinner, FormErrorMessage, VStack, Textarea } from '@chakra-ui/react'
import { yupResolver } from '@hookform/resolvers/yup'
import { getCookie } from 'cookies-next'
import { useRouter } from 'next/navigation'
import { useDropzone } from 'react-dropzone'
import { useForm, SubmitHandler } from 'react-hook-form'
import { FaUpload } from 'react-icons/fa'

import { uploadImage as UPLOAD_IMAGE, UPDATE_USER } from '@gql'
import { UserOnboardingForm } from '@types'
import { UserOnboardingSchema } from 'form/validations'

const UserOnboarding = (): JSX.Element => {
  const router = useRouter()
	const { handleSubmit, register, formState: { errors }, trigger, watch, setValue } = useForm<UserOnboardingForm>({
		resolver: yupResolver(UserOnboardingSchema)
	})

  const [uploadImage, { loading: uploadingImage }] = useMutation(UPLOAD_IMAGE)
  const [updateUser, { loading: updatingUser }] = useMutation(UPDATE_USER)

  const personImg = watch("avatar")

  const handleDrop = async (files: File[]) => {
    await uploadImage({
      variables: {
        file: files[0],
      },
      onCompleted: data => {
        setValue("avatar", data.uploadImage)
        trigger("avatar")
      },
    })
  }
  const { getInputProps, getRootProps } = useDropzone({
    accept: {
      'image/*': ['.jpg', '.jpeg', '.png', '.gif'],
    },
    disabled: uploadingImage || updatingUser,
    multiple: false,
    onDrop: handleDrop,
  })

  const onSubmit: SubmitHandler<UserOnboardingForm> = async data => {
    if (!updatingUser) { // Synchronous
      await updateUser({
        variables: {
          input: {
            id: getCookie('userId'),
            first_name: data.firstName,
            last_name: data.lastName,
            avatar: data.avatar,
            onboarding: false,
            email: "", // TODO we don't need email now?
          },
        },
        onCompleted: data => {
          router.push('/')
        },
      })
    }
  }

	return (
		<form onSubmit={handleSubmit(onSubmit)}>
      <FormControl isInvalid={Boolean(errors?.avatar)}>
        <FormLabel>Avatar</FormLabel>
        <Textarea
          as={Center}
          {...getRootProps({ className: 'dropzone' })}
          borderRadius="md"
          border="2px dashed"
          minH={{ base: "60", md: "120px" }}
          borderColor="chakra-border-color"
          h="auto"
          spacing={4}
          p={4}
          justify={personImg ? 'start' : 'center'}
          cursor="pointer"
        >
          {personImg && !uploadingImage && <Image
            src={personImg}
            loading="eager"
            maxW={{
              base: "100px",
              md: "200px"
            }}
            alt="Avatar"
          />}
          {!personImg && !uploadingImage && <Icon as={FaUpload} />}
          {uploadingImage && <Spinner size="lg" />}
        </Textarea>
        <input {...register("avatar")} {...getInputProps()} hidden />
        {((errors.avatar) != null) && <FormErrorMessage>{errors?.avatar.message}</FormErrorMessage>}
      </FormControl>
      <FormControl isInvalid={Boolean(errors?.firstName)}>
        <FormLabel>First Name</FormLabel>
        <Input
          {...register("firstName")}
        />
        {((errors.firstName) != null) && <FormErrorMessage>{errors?.firstName.message}</FormErrorMessage>}
      </FormControl>
      <FormControl isInvalid={Boolean(errors?.lastName)}>
        <FormLabel>Last Name</FormLabel>
        <Input
          {...register("lastName")}
        />
        {((errors.lastName) != null) && <FormErrorMessage>{errors?.lastName.message}</FormErrorMessage>}
      </FormControl>
      <VStack mt={5}>
        <Button isLoading={updatingUser} type="submit">Save</Button>
      </VStack>
    </form>
	)
}

export default UserOnboarding
