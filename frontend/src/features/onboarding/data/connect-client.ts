/**
 * Onboarding slice-local Connect client。
 *
 * phase07-10 §5.3：slice-local generated client 唯一承接位。
 */
import { createClient } from '@connectrpc/connect'
import { OnboardingService } from '@/gen/proto/psco/onboarding/v1/onboarding_pb'
import { transport } from '@/shared/rpc/connect-transport'

export const onboardingClient = createClient(OnboardingService, transport)