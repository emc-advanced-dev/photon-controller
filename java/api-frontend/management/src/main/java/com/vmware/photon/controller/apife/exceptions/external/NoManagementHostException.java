/*
 * Copyright 2015 VMware, Inc. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not
 * use this file except in compliance with the License.  You may obtain a copy of
 * the License at http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed
 * under the License is distributed on an "AS IS" BASIS, without warranties or
 * conditions of any kind, EITHER EXPRESS OR IMPLIED.  See the License for the
 * specific language governing permissions and limitations under the License.
 */

package com.vmware.photon.controller.apife.exceptions.external;

import com.vmware.photon.controller.api.common.exceptions.external.ErrorCode;
import com.vmware.photon.controller.api.common.exceptions.external.ExternalException;

/**
 * Gets thrown when calling deployment but no management host was configured.
 */
public class NoManagementHostException extends ExternalException {
  private final String id;

  public NoManagementHostException(String id) {
    super(ErrorCode.NO_MANAGEMENT_HOST_CONFIGURED);
    this.id = id;

    addData("id", id);
  }

  @Override
  public String getMessage() {
    return "No Management hosts were configured for Deployment #" + id;
  }
}
